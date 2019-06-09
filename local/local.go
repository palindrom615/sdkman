package local

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sdkman-cli/conf"
	"sdkman-cli/utils"
)

var e = conf.GetConf()

func IsInstalled(candidate string, version string) bool {
	target := installPath(candidate, version)
	dir, err := os.Lstat(target)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(target)
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false
}

func IsArchived(candidate string, version string) bool {
	target := archivePath(candidate, version)
	r, invalidZipFile := zip.OpenReader(target)
	defer func() {
		if r != nil {
			r.Close()
		}
	}()
	return invalidZipFile == nil
}

func Current(candidate string) (string, error) {
	p, err := os.Readlink(installPath(candidate, "current"))
	if err == nil {
		d, err := os.Stat(p)
		if err != nil {
			log.Fatal(err)
		}
		return d.Name(), nil
	}
	return "", err
}

func Installed(candidate string) ([]string, error) {
	res := []string{}
	versions, err := ioutil.ReadDir(path.Join(e.Dir, "candidates", candidate))
	if err != nil {
		return []string{}, err
	}
	for _, ver := range versions {
		res = append(res, ver.Name())
	}
	return res, nil
}

func Archive(r io.ReadCloser, candidate string, version string, completed chan<- bool) {
	f, err := os.Create(archivePath(candidate, version))
	utils.Check(err)
	println("downloading...")
	_, err = io.Copy(f, r)
	utils.Check(err)
	defer func() {
		r.Close()
		if f != nil {
			_ = f.Close()
		}
		completed <- true
	}()
}

func Unpack(candidate string, version string, archiveReady <-chan bool, installReady chan<- bool) {
	if <-archiveReady {
		println("installing...")
		if !IsArchived(candidate, version) {
			utils.Check(utils.ErrNoArchive)
		}
		_ = os.Mkdir(path.Join(e.Dir, "candidates", candidate), os.ModeDir)
		output, err := utils.Unzip(archivePath(candidate, version), os.TempDir())
		utils.Check(err)
		utils.Check(os.Rename(output[0], installPath(candidate, version)))
		installReady <- true
	}
}

func installPath(candidate string, version string) string {
	return path.Join(e.Dir, "candidates", candidate, version)
}

func archivePath(candidate string, version string) string {
	return path.Join(e.Dir, "archives", candidate+"-"+version+".zip")
}
