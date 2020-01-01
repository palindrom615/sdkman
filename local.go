package sdkman

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func arg2sdk(reg string, root string, arg string) (Sdk, error){
	sdk := strings.Split(arg, "@")
	candidate := sdk[0]
	if err := checkValidCand(root, candidate); err != nil {
		return Sdk{}, err
	}
	if len(sdk) != 2 {
		return defaultSdk(reg, root, sdk[0])
	}
	version := sdk[1]
	return Sdk{candidate, version}, nil
}

func CurrentSdks(root string) []Sdk {
	res := []Sdk{}
	for _, cand := range getCandidates(root) {
		sdk, err := CurrentSdk(root, cand)
		if err == nil {
			res = append(res, sdk)
		}
	}
	return res
}

func InstalledSdks(root string, candidate string) []Sdk {
	if versions, err := ioutil.ReadDir(candPath(root, candidate)); err == nil {
		var res []Sdk
		for _, ver := range versions {
			res = append(res, Sdk{candidate, ver.Name()})
		}
		return res
	} else {
		return []Sdk{}
	}
}

func defaultSdk(reg string, root string, candidate string) (Sdk, error) {
	if v, netErr := getDefault(reg, candidate); netErr == nil {
		return Sdk{candidate, v}, nil
	} else if curr, fsErr := CurrentSdk(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return Sdk{candidate, ""}, ErrNotOnline
	}
}

func checkValidCand(root string, candidate string) error {
	for _, can := range getCandidates(root) {
		if can == candidate {
			return nil
		}
	}
	return ErrNoCand
}

func CurrentSdk(root string, candidate string) (Sdk, error) {
	p, err := os.Readlink(Sdk{candidate, "current"}.installPath(root))
	if err == nil {
		d, _ := os.Stat(p)
		return Sdk{candidate, d.Name()}, nil
	}
	return Sdk{candidate, ""}, ErrNoCurrSdk(candidate)
}

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

func (sdk Sdk) IsArchived(root string) bool {
	return sdk.archiveFile(root) != ""
}

func (sdk Sdk) Unarchive(root string, archiveReady <-chan bool, installReady chan<- bool) error {
	if <-archiveReady {
		fmt.Printf("installing %s@%s...\n", sdk.Candidate, sdk.Version)
		if !sdk.IsArchived(root) {
			return ErrArcNotIns
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
	return ErrArcNotIns
}

func (sdk Sdk) Use(root string) error {
	return os.Symlink(sdk.installPath(root), Sdk{sdk.Candidate, "current"}.installPath(root))
}

func (sdk Sdk) checkValidVer(reg string, root string) error {
	isValid, netErr := getValidate(reg, sdk)
	if (netErr == nil && isValid) || sdk.IsInstalled(root) {
		return nil
	} else if netErr != nil {
		return ErrNotOnline
	} else {
		return ErrNoVer
	}
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
	fmt.Printf("downloading %s@%s...\n", archive.Sdk.Candidate, archive.Sdk.Version)
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
