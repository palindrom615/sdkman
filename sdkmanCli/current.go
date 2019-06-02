package sdkmanCli

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func Current(e *Env, candidate string) {
	if candidate == "" {
		installedCount := 0
		for _, c := range e.Candidates {
			CURRENT, err := currentVersion(e, c)
			if err == nil {
				fmt.Println(c + ": " + CURRENT)
				installedCount++
			}
		}
		if installedCount == 0 {
			color.Red("No candidates are in use")
		}
	} else {
		CURRENT, err := currentVersion(e, candidate)
		if err == nil {
			fmt.Println(candidate + ": " + CURRENT)
		} else {
			color.Red("Not using any version of " + candidate)
		}
	}
}
func currentVersion(e *Env, candidate string) (string, error) {
	p, err := os.Readlink(strings.Join([]string{e.CandidateDir, candidate, "current"}, string(os.PathSeparator)))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}
