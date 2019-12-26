package command

import (
	"fmt"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Current(candidate string) error {
	if candidate == "" {
		cands, vers := local.UsingCands()
		if len(cands) == 0 {
			return utils.ErrCandsNotIns
		}
		for i, _ := range cands {
			fmt.Println(cands[i] + ": " + vers[i])
		}
	} else {
		ver, err := local.UsingVer(candidate)
		if err == nil {
			fmt.Println(candidate + ": " + ver)
		} else {
			return utils.ErrCandNotIns(candidate)
		}
	}
	return nil
}
