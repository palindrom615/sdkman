package command

import (
	"fmt"
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
)

func Current(c *cli.Context) error {
	candidate := c.Args().Get(0)
	root := c.String("directory")
	if candidate == "" {
		cands, vers := local.UsingCands(root)
		if len(cands) == 0 {
			return utils.ErrCandsNotIns
		}
		for i := range cands {
			fmt.Println(cands[i] + ": " + vers[i])
		}
	} else {
		ver, err := local.UsingVer(root, candidate)
		if err == nil {
			fmt.Println(candidate + ": " + ver)
		} else {
			return utils.ErrCandNotIns(candidate)
		}
	}
	return nil
}
