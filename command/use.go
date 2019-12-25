package command

import (
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Use(candidate string, version string) error {
	utils.IsCandidateValid(candidate)
	if !local.IsInstalled(candidate, version) {
		return utils.ErrNoVersion
	}
	local.LinkCurrent(candidate, version)
	return nil
}
