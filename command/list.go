package command

import (
	"github.com/palindrom615/sdkman-cli/api"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
	"io"
)

func List(candidate string) error {
	var (
		list io.ReadCloser
		err  error
	)
	if candidate == "" {
		list, err = api.GetList()
	} else {
		utils.IsCandidateValid(candidate)
		ins, _ := local.Installed(candidate)
		curr, _ := local.Current(candidate)
		list, err = api.GetVersionsList(candidate, curr, ins)
	}
	utils.Pager(list)
	return err
}
