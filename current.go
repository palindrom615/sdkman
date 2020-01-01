package sdkman

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Current(c *cli.Context) error {
	candidate := c.Args().Get(0)
	root := c.String("directory")
	if candidate == "" {
		sdks := UsingCands(root)
		if len(sdks) == 0 {
			return ErrCandsNotIns
		}
		for _, sdk := range sdks {
			fmt.Printf("%s@%s\n", sdk.Candidate, sdk.Version)
		}
	} else {
		ver, err := UsingVer(root, candidate)
		if err == nil {
			fmt.Println(candidate + ": " + ver)
		} else {
			return ErrCandNotIns(candidate)
		}
	}
	return nil
}
