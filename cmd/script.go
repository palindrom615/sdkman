package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/scylladb/go-set/strset"
)

type envVar struct {
	name string
	val  string
}

func exportBash(paths []string, envVars []envVar) string {
	res := fmt.Sprintf("export PATH=%s:$PATH\n", strings.Join(paths, ":"))
	for _, v := range envVars {
		res += fmt.Sprintf("export %s=%s\n", v.name, v.val)
	}
	return res
}

func exportFish(paths []string, envVars []envVar) string {
	res := fmt.Sprintf("set -x PATH %s $PATH\n", strings.Join(paths, " "))
	for _, v := range envVars {
		res += fmt.Sprintf("set -x %s %s\n", v.name, v.val)
	}
	return res
}

func exportPosh(paths []string, envVars []envVar) string {
	res := fmt.Sprintf("$env:Path = \"%s;\" + $env:Path;", strings.Join(paths, ";"))
	for _, v := range envVars {
		res += fmt.Sprintf("$env:%s = \"%s\";", v.name, v.val)
	}
	return res
}

func exportWindows(paths []string, envVars []envVar) string {
	path := strset.New(strings.Split(os.Getenv("Path"), ";")...)
	path.Merge(strset.New(paths...))
	res := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"Path\", \"%s\", [System.EnvironmentVariableTarget]::User);", strings.Join(path.List(), ";"))
	for _, v := range envVars {
		res += fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"%s\", \"%s\", [System.EnvironmentVariableTarget]::User);", v.name, v.val)
	}
	return res
}
