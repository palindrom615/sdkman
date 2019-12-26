package local

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/palindrom615/sdkman-cli/conf"
	"github.com/palindrom615/sdkman-cli/store"
	"github.com/palindrom615/sdkman-cli/utils"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var e = conf.GetConf()

func candPath(candidate string) string {
	return path.Join(e.Dir, "candidates", candidate)
}
func installPath(candidate string, version string) string {
	return path.Join(candPath(candidate), version)
}

func archivePath(candidate string, version string, format string) string {
	return path.Join(e.Dir, "archives", candidate+"-"+version+"."+format)
}

func archiveFile(candidate string, version string) string {
	archives, _ := os.Open(path.Join(e.Dir, "archives"))
	arcs, _ := archives.Readdir(0)
	for _, archive := range arcs {
		if strings.HasPrefix(archive.Name(), candidate+"-"+version) {
			return path.Join(e.Dir, "archives", archive.Name())
		}
	}
	return ""
}

func IsInstalled(candidate string, version string) bool {
	target := installPath(candidate, version)
	dir, err := os.Lstat(target)
	if os.IsNotExist(err) {
		return false
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(target)
		return err == nil
	}
	return false
}

func InstalledVers(candidate string) []string {
	if versions, err := ioutil.ReadDir(candPath(candidate)); err == nil {
		var res []string
		for _, ver := range versions {
			res = append(res, ver.Name())
		}
		return res
	} else {
		return []string{}
	}
}

func UsingCands() ([]string, []string) {
	var cands, vers []string
	for _, cand := range store.GetCandidates() {
		ver, err := UsingVer(cand)
		if err == nil {
			cands = append(cands, cand)
			vers = append(vers, ver)
		}
	}
	return cands, vers
}

func IsArchived(candidate string, version string) bool {
	target := archiveFile(candidate, version)
	return target != ""
}

func Archive(r io.ReadCloser, candidate string, version string, format string, completed chan<- bool) error {
	f, err := os.Create(archivePath(candidate, version, format))
	if err != nil {
		return err
	}
	fmt.Printf("downloading %s %s...\n", candidate, version)
	_, err = io.Copy(f, r)
	if err != nil {
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

func Unarchive(candidate string, version string, archiveReady <-chan bool, installReady chan<- bool) error {
	if <-archiveReady {
		fmt.Printf("installing %s %s...\n", candidate, version)
		if !IsArchived(candidate, version) {
			return utils.ErrArcNotIns
		}
		_ = os.Mkdir(candPath(candidate), os.ModeDir|os.ModePerm)

		err := archiver.Unarchive(archiveFile(candidate, version), installPath(candidate, version))
		if err != nil {
			_ = os.RemoveAll(installPath(candidate, version))
		}
		installReady <- true
		return err
	}
	return utils.ErrArcNotIns
}

func UsingVer(candidate string) (string, error) {
	p, err := os.Readlink(installPath(candidate, "current"))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func UseVer(candidate string, version string) error {
	return os.Symlink(installPath(candidate, version), installPath(candidate, "current"))
}
