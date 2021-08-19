package cmd

import (
	"fmt"
	"strings"
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
	for i, p := range paths {
		paths[i] = strings.ReplaceAll(p, "/", "\\")
	}
	res := fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"Path\", [Environment]::GetEnvironmentVariable(\"Path\", [EnvironmentVariableTarget]::User) + \";%s\", [System.EnvironmentVariableTarget]::User);", strings.Join(paths, ";"))
	for _, v := range envVars {
		res += fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"%s\", \"%s\", [System.EnvironmentVariableTarget]::User);", v.name, v.val)
	}
	return res
}

func removeIndexes(s []string, idxs []int) []string {
	res := []string{}
	oldIdx := 0
	for _, idx := range idxs {
		res = append(res, s[oldIdx:idx]...)
		oldIdx = idx + 1
	}
	return res
}
