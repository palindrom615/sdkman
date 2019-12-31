package command

import (
	"github.com/palindrom615/sdk/api"
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	candidate := c.Args().Get(0)
	reg := c.String("registry")
	root := c.String("directory")

	if candidate == "" {
		list, err := api.GetList(reg)
		if err == nil {
			utils.Pager(list)
		}
		return err
	} else {
		if err := utils.CheckValidCand(root, candidate); err != nil {
			return err
		}
		ins := local.InstalledVers(root, candidate)
		curr, _ := local.UsingVer(root, candidate)
		list, err := api.GetVersionsList(reg, candidate, curr, ins)
		utils.Pager(list)
		return err
	}
}
