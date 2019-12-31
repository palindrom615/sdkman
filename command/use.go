package command

import (
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
)

func Use(c *cli.Context) error {
	candidate := c.Args().Get(0)
	version := c.Args().Get(1)
	root := c.String("directory")
	if err := utils.CheckValidCand(root, candidate); err != nil {
		return err
	}
	if !local.IsInstalled(root, candidate, version) {
		return utils.ErrVerNotIns
	}
	return local.UseVer(root, candidate, version)
}
