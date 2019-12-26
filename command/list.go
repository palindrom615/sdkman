package command

import (
	"github.com/palindrom615/sdkman-cli/api"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func List(candidate string) error {
	if candidate == "" {
		list, err := api.GetList()
		if err == nil {
			utils.Pager(list)
		}
		return err
	} else {
		if err := utils.CheckValidCand(candidate); err != nil {
			return err
		}
		ins := local.InstalledVers(candidate)
		curr, _ := local.UsingVer(candidate)
		list, err := api.GetVersionsList(candidate, curr, ins)
		utils.Pager(list)
		return err
	}
}
