package command

import (
	"fmt"
	"sdkman-cli/local"
	"sdkman-cli/store"

	"github.com/fatih/color"
)

func Current(candidate string) {
	if candidate == "" {
		installedCount := 0
		for _, c := range store.GetCandidates() {
			if printCurrent(c) == nil {
				installedCount++
			}
		}
		if installedCount == 0 {
			color.Red("No candidates are in use")
		}
	} else {
		if printCurrent(candidate) != nil {
			color.Red("Not using any version of " + candidate)
		}
	}
}

func printCurrent(c string) error {
	ver, err := local.CurrentVersion(c)
	if err == nil {
		fmt.Println(c + ": " + ver)
	}
	return err
}
