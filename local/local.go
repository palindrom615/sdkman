package local

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/palindrom615/sdk/store"
	"github.com/palindrom615/sdk/utils"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func candPath(root string, candidate string) string {
	return path.Join(root, "candidates", candidate)
}
func installPath(root string, candidate string, version string) string {
	return path.Join(candPath(root, candidate), version)
}

func archivePath(root string, candidate string, version string, format string) string {
	return path.Join(root, "archives", candidate+"-"+version+"."+format)
}

func archiveFile(root string, candidate string, version string) string {
	archives, _ := os.Open(path.Join(root, "archives"))
	arcs, _ := archives.Readdir(0)
	for _, archive := range arcs {
		if strings.HasPrefix(archive.Name(), candidate+"-"+version) {
			return path.Join(root, "archives", archive.Name())
		}
	}
	return ""
}

func MkdirIfNotExist(root string) error {
	candDir := path.Join(root, "candidates")
	arcDir := path.Join(root, "archives")
	e := os.MkdirAll(candDir, os.ModeDir|os.ModePerm)
	if e != nil {
		return e
	}
	return os.MkdirAll(arcDir, os.ModeDir|os.ModePerm)
}

func IsInstalled(root string, candidate string, version string) bool {
	dir, err := os.Lstat(installPath(root, candidate, version))
	if err != nil {
		return false
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(installPath(root, candidate, version))
		return err == nil
	}
	return false
}

func InstalledVers(root string, candidate string) []string {
	if versions, err := ioutil.ReadDir(candPath(root, candidate)); err == nil {
		var res []string
		for _, ver := range versions {
			res = append(res, ver.Name())
		}
		return res
	} else {
		return []string{}
	}
}

func UsingCands(root string) ([]string, []string) {
	var cands, vers []string
	for _, cand := range store.GetCandidates(root) {
		ver, err := UsingVer(root, cand)
		if err == nil {
			cands = append(cands, cand)
			vers = append(vers, ver)
		}
	}
	return cands, vers
}

func IsArchived(root string, candidate string, version string) bool {
	return archiveFile(root, candidate, version) != ""
}

func Archive(r io.ReadCloser, root string, candidate string, version string, format string, completed chan<- bool) error {
	if format == "gz" {
		format = "tar.gz"
	}
	f, err := os.Create(archivePath(root, candidate, version, format))
	if err != nil {
		completed <- false
		return err
	}
	fmt.Printf("downloading %s %s...\n", candidate, version)
	_, err = io.Copy(f, r)
	if os.IsNotExist(err) {
		completed <- false
		return err
	}
	defer func() {
		r.Close()
		if f != nil {
			_ = f.Close()
		}
		completed <- true
	}()
	return nil
}

func Unarchive(root string, candidate string, version string, archiveReady <-chan bool, installReady chan<- bool) error {
	if <-archiveReady {
		fmt.Printf("installing %s %s...\n", candidate, version)
		if !IsArchived(root, candidate, version) {
			return utils.ErrArcNotIns
		}
		_ = os.Mkdir(candPath(root, candidate), os.ModeDir|os.ModePerm)

		wd := installPath(root, candidate, version)
		err := archiver.Unarchive(archiveFile(root, candidate, version), wd)
		if err != nil {
			installReady <- false
			_ = os.RemoveAll(wd)
		}

		// for nested directory like java:
		if l, _ := ioutil.ReadDir(wd); len(l) == 1 && l[0].IsDir() {
			nestedRoot := l[0].Name()
			inside, _ := ioutil.ReadDir(path.Join(wd, nestedRoot))
			for _, c := range inside {
				os.Rename(path.Join(wd, nestedRoot, c.Name()), path.Join(wd, c.Name()))
			}
			os.Remove(nestedRoot)
		}

		installReady <- true
		return err
	}
	return utils.ErrArcNotIns
}

func UsingVer(root string, candidate string) (string, error) {
	p, err := os.Readlink(installPath(root, candidate, "current"))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func UseVer(root string, candidate string, version string) error {
	return os.Symlink(installPath(root, candidate, version), installPath(root, candidate, "current"))
}
