package local

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/palindrom615/sdkman-cli/conf"
	"github.com/palindrom615/sdkman-cli/utils"
	"io/ioutil"
	"os"
	"path"
)

var e = conf.GetConf()

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

func Installed(candidate string) []string {
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

func Unpack(candidate string, version string, archiveReady <-chan bool, installReady chan<- bool) error {
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

func candPath(candidate string) string {
	return path.Join(e.Dir, "candidates", candidate)
}
func installPath(candidate string, version string) string {
	return path.Join(candPath(candidate), version)
}
