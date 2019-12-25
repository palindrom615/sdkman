package command

import (
	"fmt"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/store"

	"github.com/fatih/color"
)

func Current(candidate string) error {
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
	return nil
}

func printCurrent(c string) error {
	ver, err := local.Current(c)
	if err == nil {
		fmt.Println(c + ": " + ver)
	}
	return err
}
