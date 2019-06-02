package sdkmanCli

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path"
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
	p, err := os.Readlink(path.Join(e.CandidateDir, candidate, "current"))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func installed(e *Env, candidate string) ([]string, error) {
	res := []string{}
	vers, err := ioutil.ReadDir(path.Join(e.CandidateDir, candidate))
	for _, ver := range vers {
		res = append(res, ver.Name())
	}
	return res, err
}
