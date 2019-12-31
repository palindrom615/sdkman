package command

import (
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
)

func Use(candidate string, version string) error {
	if err := utils.CheckValidCand(candidate); err != nil {
		return err
	}
	if !local.IsInstalled(candidate, version) {
		return utils.ErrVerNotIns
	}
	return local.UseVer(candidate, version)
}
