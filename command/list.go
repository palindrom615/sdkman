package command

import (
	"sdkman-cli/api"
	"sdkman-cli/local"
	"sdkman-cli/utils"
)

func List(candidate string) {

	if candidate == "" {
		list, err := api.GetList()
		if err != nil {
			utils.Check(utils.ErrNotOnline)
		}
		utils.Pager(list)
	} else {
		ins, _ := local.Installed(candidate)
		curr, _ := local.Current(candidate)
		list, err := api.GetVersionsList(candidate, curr, ins)
		if err != nil {
			utils.Check(utils.ErrNotOnline)
		}
		utils.Pager(list)
	}
}
