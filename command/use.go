package command

import (
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Use(candidate string, version string) error {
	if err := utils.CheckValidCand(candidate); err != nil {
		return err
	}
	if !local.IsInstalled(candidate, version) {
		return utils.ErrVerNotIns
	}
	return local.SetCurrVer(candidate, version)
}
