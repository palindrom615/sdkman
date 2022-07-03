package sdk

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/otiai10/copy"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/store"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Sdk represents each version of sdk
// ex) Sdk{Candidate: "java", Version: "8.0.232-zulu"}
type Sdk struct {
	Candidate string
	Version   string
}

func candPath(root string, candidate string) string {
	return path.Join(root, "candidates", candidate)
}

func (sdk Sdk) installPath(root string) string {
	return path.Join(candPath(root, sdk.Candidate), sdk.Version)
}

func (sdk Sdk) archiveFile(root string) string {
	archives, _ := os.Open(path.Join(root, "archives"))
	arcs, _ := archives.Readdir(0)
	for _, archive := range arcs {
		if strings.HasPrefix(archive.Name(), sdk.Candidate+"-"+sdk.Version) {
			return path.Join(root, "archives", archive.Name())
		}
	}
	return ""
}

// IsInstalled returns if sdk is installed or not
func (sdk Sdk) IsInstalled(root string) bool {
	dir, err := os.Lstat(sdk.installPath(root))
	if err != nil {
		return false
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(sdk.installPath(root))
		return err == nil
	}
	return false
}

// IsArchived returns if archive file of sdk exists or not
func (sdk Sdk) IsArchived(root string) bool {
	return sdk.archiveFile(root) != ""
}

// Unarchive extracts archive file of sdk into "candidates" directory
func (sdk Sdk) Unarchive(root string, archiveReady <-chan bool, installReady chan<- bool) error {
	if <-archiveReady {
		fmt.Printf("installing %s@%s...\n", sdk.Candidate, sdk.Version)
		if !sdk.IsArchived(root) {
			return errors.ErrArcNotIns
		}
		_ = os.Mkdir(candPath(root, sdk.Candidate), os.ModeDir|os.ModePerm)

		wd := sdk.installPath(root)
		err := archiver.Unarchive(sdk.archiveFile(root), wd)
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
	return errors.ErrArcNotIns
}

// Use links sdk with symlink named "current" so the sdk is used as default
func (sdk Sdk) Use(root string) error {
	os.Remove(Sdk{sdk.Candidate, "current"}.installPath(root))
	err := os.Symlink(sdk.installPath(root), Sdk{sdk.Candidate, "current"}.installPath(root))
	if err != nil {
		// windows requires admin privilege to make symlink and I don't want to
		copy.Copy(sdk.installPath(root), Sdk{sdk.Candidate, "current"}.installPath(root))
	}
	return nil
}

func (sdk Sdk) CheckValidVer(reg string, root string) error {
	isValid, netErr := pkgs.GetValidate(reg, sdk)
	if (netErr == nil && isValid) || sdk.IsInstalled(root) {
		return nil
	} else if netErr != nil {
		return errors.ErrNotOnline
	} else {
		return errors.ErrNoVer
	}
}

// GetFromVersionString versionString format: e.g. "kotlin@1.7.0", "java@17.0.3-tem"
func GetFromVersionString(registry string, sdkHome string, versionString string) (Sdk, error) {
	sdk := strings.Split(versionString, "@")
	candidate := sdk[0]
	if err := pkgs.CheckValidCand(sdkHome, candidate); err != nil {
		return Sdk{}, err
	}
	if len(sdk) != 2 {
		return DefaultSdk(registry, sdkHome, sdk[0])
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
	if _, err := os.Stat(Sdk{candidate, "current"}.installPath(root)); err != nil {
		return Sdk{candidate, ""}, errors.ErrNoCurrSdk(candidate)
	}
	p, err := os.Readlink(Sdk{candidate, "current"}.installPath(root))
	if err == nil {
		d, _ := os.Stat(p)
		return Sdk{candidate, d.Name()}, nil
	}
	// if directory 'current' is not symlink
	return Sdk{candidate, "current"}, nil
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
	if v, netErr := pkgs.GetDefault(reg, candidate); netErr == nil {
		return Sdk{candidate, v}, nil
	} else if curr, fsErr := CurrentSdk(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return Sdk{candidate, ""}, errors.ErrNotOnline
	}
}
