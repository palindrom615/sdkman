package local

import (
	"archive/zip"
	"github.oom/palindrom615/sdkman-cli/utils"
	"io"
	"os"
	"path"
)

func IsArchived(candidate string, version string) bool {
	target := archivePath(candidate, version)
	r, invalidZipFile := zip.OpenReader(target)
	defer func() {
		if r != nil {
			r.Close()
		}
	}()
	return invalidZipFile == nil
}

func Archive(r io.ReadCloser, candidate string, version string, completed chan<- bool) {
	f, err := os.Create(archivePath(candidate, version))
	utils.Check(err)
	println("downloading...")
	_, err = io.Copy(f, r)
	utils.Check(err)
	defer func() {
		r.Close()
		if f != nil {
			_ = f.Close()
		}
		completed <- true
	}()
}

func archivePath(candidate string, version string) string {
	return path.Join(e.Dir, "archives", candidate+"-"+version+".zip")
}
