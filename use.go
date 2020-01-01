package sdkman

import (
	"github.com/urfave/cli/v2"
)

func Use(c *cli.Context) error {
	candidate := c.Args().Get(0)
	version := c.Args().Get(1)
	sdk := Sdk{candidate, version}
	root := c.String("directory")
	if err := checkValidCand(root, candidate); err != nil {
		return err
	}
	if !sdk.IsInstalled(root) {
		return ErrVerNotIns
	}
	return sdk.UseVer(root)
}
