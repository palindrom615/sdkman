package sdkman

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"path"
	"strings"
)

type envVar struct {
	name string
	val  string
}

func Export(c *cli.Context) error {
	shell := c.Args().Get(0)
	root := c.String("directory")
	sdks := UsingCands(c.String("directory"))
	if len(sdks) == 0 {
		fmt.Println("")
		return nil
	}
	paths := []string{}
	homes := []envVar{}
	for _, sdk := range sdks {
		candHome := path.Join(root, "candidates", sdk.Candidate, "current")
		paths = append(paths, path.Join(candHome, "bin"))
		homes = append(homes, envVar{fmt.Sprintf("%s_HOME", strings.ToUpper(sdk.Candidate)), candHome})
	}

	if shell == "bash" || shell == "" {
		evalBash(paths, homes)
	}
	return nil
}

func evalBash(paths []string, envVars []envVar) {
	fmt.Println("export PATH=" + strings.Join(paths, ":") + ":$PATH")
	for _, v := range envVars {
		fmt.Println("export " + v.name + "=" + v.val)
	}
}
