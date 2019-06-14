package command

import (
	"io"
	"github.oom/palindrom615/sdkman-cli/api"
	"github.oom/palindrom615/sdkman-cli/local"
	"github.oom/palindrom615/sdkman-cli/utils"
)

func List(candidate string) {
	var (
		list io.ReadCloser
		err  error
	)
	if candidate == "" {
		list, err = api.GetList()
	} else {
		utils.CheckValid(candidate)
		ins, _ := local.Installed(candidate)
		curr, _ := local.Current(candidate)
		list, err = api.GetVersionsList(candidate, curr, ins)
	}
	utils.Check(err)
	utils.Pager(list)
}
