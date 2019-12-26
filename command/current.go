package command

import (
	"fmt"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/store"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Current(candidate string) error {
	if candidate == "" {
		installedCount := 0
		for _, c := range store.GetCandidates() {
			if printCurrVer(c) == nil {
				installedCount++
			}
		}
		if installedCount == 0 {
			return utils.ErrCandsNotIns
		}
	} else {
		if printCurrVer(candidate) != nil {
			return utils.ErrCandNotIns(candidate)
		}
	}
	return nil
}

func printCurrVer(c string) error {
	ver, err := local.GetCurrVer(c)
	if err == nil {
		fmt.Println(c + ": " + ver)
	}
	return err
}
