package local

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sdkman-cli/conf"
)

var e = conf.GetConf()

func IsVersionExists(candidate string, version string ) bool {
	target := path.Join(e.Dir, "candidates", candidate, version)
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

func IsArchiveExists(candidate string, version string) bool {
	target := path.Join(e.Dir, "archives",  candidate + "-" + version + ".zip")
	_, err := os.Stat(target)
	return os.IsNotExist(err)
}

func CurrentVersion(candidate string) (string, error) {
	p, err := os.Readlink(path.Join(e.Dir, "candidates", candidate, "current"))
	if err == nil {
		d, err := os.Stat(p)
		if err != nil {
			log.Fatal(err)
		}
		return d.Name(), nil
	}
	return "", err
}

func InstalledVersions(candidate string) ([]string, error) {
	res := []string{}
	versions, err := ioutil.ReadDir(path.Join(e.Dir, "candidates", candidate))
	for _, ver := range versions {
		res = append(res, ver.Name())
	}
	return res, err
}
