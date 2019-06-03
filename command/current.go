package command

import (
	"fmt"
	"sdkman-cli/store"

	"github.com/fatih/color"
)

func Current(candidate string) {
	if candidate == "" {
		installedCount := 0
		for _, c := range store.GetCandidates() {
			CURRENT, err := currentVersion(c)
			if err == nil {
				fmt.Println(c + ": " + CURRENT)
				installedCount++
			}
		}
		if installedCount == 0 {
			color.Red("No candidates are in use")
		}
	} else {
		CURRENT, err := currentVersion(candidate)
		if err == nil {
			fmt.Println(candidate + ": " + CURRENT)
		} else {
			color.Red("Not using any version of " + candidate)
		}
	}
}
