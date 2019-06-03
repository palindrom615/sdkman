package command

import (
	"sdkman-cli/api"
	"sdkman-cli/local"
)

func List(candidate string) {

	if candidate == "" {
		Pager(string(api.GetList()))
	} else {
		ins, _ := local.InstalledVersions(candidate)
		curr, _ := local.CurrentVersion(candidate)
		Pager(string(api.GetVersionsList(candidate, curr, ins)))
	}
}
