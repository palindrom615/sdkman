package pkgs

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/palindrom615/sdkman/errors"
	"github.com/otiai10/copy"
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
	if err != nil{
		// windows requires admin privilege to make symlink and I don't want to
		copy.Copy(sdk.installPath(root), Sdk{sdk.Candidate, "current"}.installPath(root))
	}
	return nil
}

func (sdk Sdk) CheckValidVer(reg string, root string) error {
	isValid, netErr := GetValidate(reg, sdk)
	if (netErr == nil && isValid) || sdk.IsInstalled(root) {
		return nil
	} else if netErr != nil {
		return errors.ErrNotOnline
	} else {
		return errors.ErrNoVer
	}
}
