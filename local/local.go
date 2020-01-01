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

type Sdk struct {
	Candidate string
	Version   string
}
type Archive struct {
	Sdk    Sdk
	Format string
}


func candPath(root string, candidate string) string {
	return path.Join(root, "candidates", candidate)
}
func (sdk Sdk) installPath(root string) string {
	return path.Join(candPath(root, sdk.Candidate), sdk.Version)
}

func (archive Archive) archivePath(root string) string {
	fileName := fmt.Sprintf("%s-%s.%s", archive.Sdk.Candidate, archive.Sdk.Version, archive.Format)
	return path.Join(root, "archives", fileName)
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

func MkdirIfNotExist(root string) error {
	candDir := path.Join(root, "candidates")
	arcDir := path.Join(root, "archives")
	e := os.MkdirAll(candDir, os.ModeDir|os.ModePerm)
	if e != nil {
		return e
	}
	return os.MkdirAll(arcDir, os.ModeDir|os.ModePerm)
}

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

func UsingCands(root string) []Sdk {
	res := []Sdk{}
	for _, cand := range store.GetCandidates(root) {
		ver, err := UsingVer(root, cand)
		if err == nil {
			res = append(res, Sdk{cand, ver})
		}
	}
	return res
}

func (sdk Sdk) IsArchived(root string) bool {
	return sdk.archiveFile(root) != ""
}

func (archive Archive) Save(r io.ReadCloser, root string, completed chan<- bool) error {
	if archive.Format == "gz" || archive.Format == "bz2" || archive.Format == "xz" {
		archive.Format = "tar." + archive.Format
	}
	f, err := os.Create(archive.archivePath(root))
	if err != nil {
		completed <- false
		return err
	}
	fmt.Printf("downloading %s %s...\n", archive.Sdk.Candidate, archive.Sdk.Version)
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

func (sdk Sdk) Unarchive(root string, archiveReady <-chan bool, installReady chan<- bool) error {
	if <-archiveReady {
		fmt.Printf("installing %s %s...\n", sdk.Candidate, sdk.Version)
		if !sdk.IsArchived(root) {
			return utils.ErrArcNotIns
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
	return utils.ErrArcNotIns
}

func UsingVer(root string, candidate string) (string, error) {
	p, err := os.Readlink(Sdk{candidate, "current"}.installPath(root))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func (sdk Sdk) UseVer(root string) error {
	return os.Symlink(sdk.installPath(root), Sdk{sdk.Candidate, "current"}.installPath(root))
}
