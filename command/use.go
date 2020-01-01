package command

import (
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
)

func Use(c *cli.Context) error {
	candidate := c.Args().Get(0)
	version := c.Args().Get(1)
	sdk := local.Sdk{candidate, version}
	root := c.String("directory")
	if err := utils.CheckValidCand(root, candidate); err != nil {
		return err
	}
	if !sdk.IsInstalled(root) {
		return utils.ErrVerNotIns
	}
	return sdk.UseVer(root)
}
