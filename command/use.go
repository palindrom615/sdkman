package command

import (
	"github.oom/palindrom615/sdkman-cli/local"
	"github.oom/palindrom615/sdkman-cli/utils"
)

func Use(candidate string, version string) {
	utils.CheckValid(candidate)
	if !local.IsInstalled(candidate, version) {
		utils.Check(utils.ErrNoVersion)
	}
	local.LinkCurrent(candidate, version)
}
