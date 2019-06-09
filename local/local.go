package local

import (
	"io/ioutil"
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
	} else {
		utils.Check(err)
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(target)
		utils.Check(err)
		return true
	}
	return false
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

func Unpack(candidate string, version string, archiveReady <-chan bool, installReady chan<- bool) {
	if <-archiveReady {
		println("installing...")
		if !IsArchived(candidate, version) {
			utils.Check(utils.ErrNoArchive)
		}
		_ = os.Mkdir(path.Join(e.Dir, "candidates", candidate), os.ModeDir)

		tmpDir := path.Join(os.TempDir(), candidate+"-"+version)
		defer os.RemoveAll(tmpDir)
		_, err := utils.Unzip(archivePath(candidate, version), tmpDir)
		utils.Check(err)
		res, _ := ioutil.ReadDir(tmpDir)
		result := path.Join(tmpDir, res[0].Name())
		utils.Check(os.Rename(result, installPath(candidate, version)))
		installReady <- true
	}
}

func installPath(candidate string, version string) string {
	return path.Join(e.Dir, "candidates", candidate, version)
}
