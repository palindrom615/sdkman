package sdkman

import (
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	candidate := c.Args().Get(0)
	reg := c.String("registry")
	root := c.String("directory")

	if candidate == "" {
		list, err := getList(reg)
		if err == nil {
			pager(list)
		}
		return err
	} else {
		if err := checkValidCand(root, candidate); err != nil {
			return err
		}
		ins := InstalledVers(root, candidate)
		curr, _ := UsingVer(root, candidate)
		list, err := getVersionsList(reg, candidate, curr, ins)
		pager(list)
		return err
	}
}
