package pkgs

import (
	"fmt"
	"github.com/palindrom615/sdkman/sdk"
	"io"
	"os"
	"path"
)

// Archive represents downloaded compressed files in "archives" directory
type Archive struct {
	Sdk    sdk.Sdk
	Format string
}

func (archive Archive) archivePath(root string) string {
	fileName := fmt.Sprintf("%s-%s.%s", archive.Sdk.Candidate, archive.Sdk.Version, archive.Format)
	return path.Join(root, "archives", fileName)
}

// Save saves bytes read from ReadCloser channel into archive file
func (archive Archive) Save(r io.ReadCloser, root string, completed chan<- bool) error {
	f, err := os.Create(archive.archivePath(root))
	if err != nil {
		completed <- false
		return err
	}
	fmt.Printf("saving %s@%s...\n", archive.Sdk.Candidate, archive.Sdk.Version)
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
