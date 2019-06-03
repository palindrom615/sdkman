package command

import (
	"sdkman-cli/api"
)

func List(candidate string) {

	if candidate == "" {
		Pager(string(api.GetList()))
	} else {
		ins, _ := installed(candidate)
		curr, _ := currentVersion(candidate)
		Pager(string(api.GetVersionsList(candidate, curr, ins)))
	}
}
