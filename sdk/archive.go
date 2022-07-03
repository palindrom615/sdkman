package sdk

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Archive represents downloaded compressed files in "archives" directory
type Archive struct {
	Sdk     Sdk
	Format  string
	SdkHome string
}

func (archive Archive) archivePath(root string) string {
	fileName := fmt.Sprintf("%s-%s.%s", archive.Sdk.Candidate, archive.Sdk.Version, archive.Format)
	return filepath.Join(root, "archives", fileName)
}

// Save saves bytes read from ReadCloser channel into archive file
func (archive Archive) Save(r io.ReadCloser) error {
	f, err := os.Create(archive.archivePath(archive.SdkHome))
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	if os.IsNotExist(err) {
		return err
	}
	defer func() {
		r.Close()
		if f != nil {
			_ = f.Close()
		}
	}()
	return nil
}
