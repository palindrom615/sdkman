package local

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

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
