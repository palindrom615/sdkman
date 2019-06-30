package local

import (
	"github.com/palindrom615/sdkman-cli/utils"
	"os"
)

func Current(candidate string) (string, error) {
	p, err := os.Readlink(installPath(candidate, "current"))
	if err == nil {
		d, err := os.Stat(p)
		utils.Check(err)
		return d.Name(), nil
	}
	return "", err
}

func LinkCurrent(candidate string, version string) {
	utils.Check(os.Symlink(installPath(candidate, version), installPath(candidate, "current")))
}
