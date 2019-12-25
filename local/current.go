package local

import (
	"os"
)

func Current(candidate string) (string, error) {
	p, err := os.Readlink(installPath(candidate, "current"))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func LinkCurrent(candidate string, version string) error {
	return os.Symlink(installPath(candidate, version), installPath(candidate, "current"))
}
