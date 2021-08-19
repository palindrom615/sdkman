package pkgs

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/store"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

// MkdirIfNotExist creates "candidates" and "archives" directory
func MkdirIfNotExist(root string) error {
	candDir := path.Join(root, "candidates")
	arcDir := path.Join(root, "archives")
	e := os.MkdirAll(candDir, os.ModeDir|os.ModePerm)
	if e != nil {
		return e
	}
	return os.MkdirAll(arcDir, os.ModeDir|os.ModePerm)
}

func Pager(pages io.ReadCloser) {
	pager := os.Getenv("PAGER")

	if pager == "" {
		if runtime.GOOS == "windows" {
			pager = "more"
		} else {
			pager = "less"
		}
	}
	c1 := exec.Command(pager)
	c1.Stdin = pages
	c1.Stdout = os.Stdout
	_ = c1.Start()
	_ = c1.Wait()
	defer pages.Close()
}

func Arg2sdk(reg string, root string, arg string) (Sdk, error) {
	sdk := strings.Split(arg, "@")
	candidate := sdk[0]
	if err := CheckValidCand(root, candidate); err != nil {
		return Sdk{}, err
	}
	if len(sdk) != 2 {
		return DefaultSdk(reg, root, sdk[0])
	}
	version := sdk[1]
	return Sdk{candidate, version}, nil
}

// CurrentSdks returns every Sdk that is linked via "current"
func CurrentSdks(root string) []Sdk {
	res := []Sdk{}
	for _, cand := range store.GetCandidates(root) {
		sdk, err := CurrentSdk(root, cand)
		if err == nil {
			res = append(res, sdk)
		}
	}
	return res
}

// CurrentSdk returns sdk of specified candidate which is linked with "current"
func CurrentSdk(root string, candidate string) (Sdk, error) {
	p, err := os.Readlink(Sdk{candidate, "current"}.installPath(root))
	if err == nil {
		d, _ := os.Stat(p)
		return Sdk{candidate, d.Name()}, nil
	}
	return Sdk{candidate, ""}, errors.ErrNoCurrSdk(candidate)
}

// InstalledSdks returns every installed Sdk of specified candidate
func InstalledSdks(root string, candidate string) []Sdk {
	versions, err := ioutil.ReadDir(candPath(root, candidate))
	if err != nil {
		return []Sdk{}
	}
	var res []Sdk
	for _, ver := range versions {
		res = append(res, Sdk{candidate, ver.Name()})
	}
	return res
}

func DefaultSdk(reg string, root string, candidate string) (Sdk, error) {
	if v, netErr := GetDefault(reg, candidate); netErr == nil {
		return Sdk{candidate, v}, nil
	} else if curr, fsErr := CurrentSdk(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return Sdk{candidate, ""}, errors.ErrNotOnline
	}
}

func CheckValidCand(root string, candidate string) error {
	for _, can := range store.GetCandidates(root) {
		if can == candidate {
			return nil
		}
	}
	return errors.ErrNoCand
}
