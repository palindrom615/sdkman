package sdk

import (
	"github.com/mholt/archiver/v3"
	"github.com/otiai10/copy"
	"github.com/palindrom615/sdkman/api"
	"github.com/palindrom615/sdkman/custom_errors"
	"github.com/palindrom615/sdkman/store"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Sdk represents each version of sdk
// ex) Sdk{Candidate: "java", Version: "8.0.232-zulu"}
type Sdk struct {
	Candidate string
	Version   string
	SdkHome   string
}

func candPath(root string, candidate string) string {
	return filepath.Join(root, "candidates", candidate)
}

func (sdk Sdk) installPath() string {
	return filepath.Join(candPath(sdk.SdkHome, sdk.Candidate), sdk.Version)
}

func (sdk Sdk) archiveFile() string {
	archives, _ := os.Open(filepath.Join(sdk.SdkHome, "archives"))
	arcs, _ := archives.Readdir(0)
	for _, archive := range arcs {
		if strings.HasPrefix(archive.Name(), sdk.Candidate+"-"+sdk.Version) {
			return filepath.Join(sdk.SdkHome, "archives", archive.Name())
		}
	}
	return ""
}

// IsInstalled returns if sdk is installed or not
func (sdk Sdk) IsInstalled() bool {
	dir, err := os.Lstat(sdk.installPath())
	if err != nil {
		return false
	}
	mode := dir.Mode()
	if mode.IsDir() {
		return true
	} else if mode&os.ModeSymlink != 0 {
		_, err := os.Readlink(sdk.installPath())
		return err == nil
	}
	return false
}

// IsArchived returns if archive file of sdk exists or not
func (sdk Sdk) IsArchived() bool {
	return sdk.archiveFile() != ""
}

// Unarchive extracts archive file of sdk into "candidates" directory
func (sdk Sdk) Unarchive() error {
	if !sdk.IsArchived() {
		return custom_errors.ErrArcNotIns
	}
	_ = os.Mkdir(candPath(sdk.SdkHome, sdk.Candidate), os.ModeDir|os.ModePerm)

	wd := sdk.installPath()
	err := archiver.Unarchive(sdk.archiveFile(), wd)
	if err != nil {
		_ = os.RemoveAll(wd)
	}

	// for nested directory like java:
	if l, _ := ioutil.ReadDir(wd); len(l) == 1 && l[0].IsDir() {
		nestedRoot := l[0].Name()
		inside, _ := ioutil.ReadDir(filepath.Join(wd, nestedRoot))
		for _, c := range inside {
			os.Rename(filepath.Join(wd, nestedRoot, c.Name()), filepath.Join(wd, c.Name()))
		}
		os.Remove(nestedRoot)
	}
	return err
}

// Use links sdk with symlink named "current" so the sdk is used as default
func (sdk Sdk) Use() error {
	os.Remove(Sdk{sdk.Candidate, "current", sdk.SdkHome}.installPath())

	targetPath, _ := filepath.Abs(sdk.installPath())
	err := os.Symlink(targetPath, Sdk{sdk.Candidate, "current", sdk.SdkHome}.installPath())
	if err != nil {
		// windows requires admin privilege to make symlink and I don't want to
		copy.Copy(sdk.installPath(), Sdk{sdk.Candidate, "current", sdk.SdkHome}.installPath())
	}
	return nil
}

func (sdk Sdk) Install(registry string) error {
	if sdk.IsInstalled() {
		return custom_errors.ErrVerExists
	}
	if sdk.IsArchived() {
		println(sdk.ToString() + ": use cached binary " + sdk.archiveFile())
	} else {
		err := sdk.Download(registry)
		if err != nil {
			return err
		}
	}
	err := sdk.Unarchive()
	if err != nil {
		return err
	}
	println(sdk.ToString() + ": installed ")
	return nil
}

func (sdk Sdk) Download(registry string) error {
	if err := sdk.CheckValidVer(registry); err != nil {
		return err
	}
	s, t, err := api.GetDownload(registry, sdk.Candidate, sdk.Version)
	if err != nil {
		return err
	}
	archive := Archive{sdk, t, sdk.SdkHome}
	err = archive.Save(s)
	if err != nil {
		return err
	}
	println(sdk.ToString() + ": downloaded")
	return nil
}

func (sdk Sdk) CheckValidVer(reg string) error {
	isValid, netErr := api.GetValidate(reg, sdk.Candidate, sdk.Version)
	if (netErr == nil && isValid) || sdk.IsInstalled() {
		return nil
	} else if netErr != nil {
		return custom_errors.ErrNotOnline
	} else {
		return custom_errors.ErrNoVer
	}
}

func (sdk Sdk) ToString() string {
	return sdk.Candidate + "@" + sdk.Version
}

// GetFromVersionString versionString format: e.g. "kotlin@1.7.0", "java@17.0.3-tem"
func GetFromVersionString(registry string, sdkHome string, versionString string) (Sdk, error) {
	sdk := strings.Split(versionString, "@")
	candidate := sdk[0]
	s := store.Store{sdkHome}
	if !s.HasCandidate(candidate) {
		return Sdk{}, custom_errors.ErrNoCand
	}
	if len(sdk) != 2 {
		return DefaultSdk(registry, sdkHome, sdk[0])
	}
	version := sdk[1]
	return Sdk{candidate, version, sdkHome}, nil
}

// CurrentSdks returns every Sdk that is linked via "current"
func CurrentSdks(root string) []Sdk {
	res := []Sdk{}
	s := store.Store{root}
	for _, cand := range s.GetCandidates() {
		sdk, err := CurrentSdk(root, cand)
		if err == nil {
			res = append(res, sdk)
		}
	}
	return res
}

// CurrentSdk returns sdk of specified candidate which is linked with "current"
func CurrentSdk(root string, candidate string) (Sdk, error) {
	if _, err := os.Stat(Sdk{candidate, "current", root}.installPath()); err != nil {
		return Sdk{candidate, "", root}, custom_errors.ErrNoCurrSdk(candidate)
	}
	p, err := os.Readlink(Sdk{candidate, "current", root}.installPath())
	if err == nil {
		d, _ := os.Stat(p)
		return Sdk{candidate, d.Name(), root}, nil
	}
	// if directory 'current' is not symlink
	return Sdk{candidate, "current", root}, nil
}

// InstalledSdks returns every installed Sdk of specified candidate
func InstalledSdks(root string, candidate string) []Sdk {
	versions, err := ioutil.ReadDir(candPath(root, candidate))
	if err != nil {
		return []Sdk{}
	}
	var res []Sdk
	for _, ver := range versions {
		res = append(res, Sdk{candidate, ver.Name(), root})
	}
	return res
}

func DefaultSdk(reg string, root string, candidate string) (Sdk, error) {
	if v, netErr := api.GetDefault(reg, candidate); netErr == nil {
		return Sdk{candidate, v, root}, nil
	} else if curr, fsErr := CurrentSdk(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return Sdk{candidate, "", root}, custom_errors.ErrNotOnline
	}
}
