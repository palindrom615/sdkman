package command

import (
	"fmt"
	"github.com/palindrom615/sdk/conf"
	"github.com/palindrom615/sdk/local"
	"path"
	"strings"
)

type envVar struct {
	name string
	val  string
}

func Export(shell string) error {
	e := conf.GetConf()
	cands, _ := local.UsingCands()
	if len(cands) == 0 {
		fmt.Println("")
		return nil
	}
	paths := []string{}
	homes := []envVar{}
	for _, cand := range cands {
		candHome := path.Join(e.Dir, "candidates", cand, "current")
		paths = append(paths, path.Join(candHome, "bin"))
		homes = append(homes, envVar{fmt.Sprintf("%s_HOME", strings.ToUpper(cand)), candHome})
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
