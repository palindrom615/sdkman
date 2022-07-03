package script

import (
	"fmt"
	"github.com/palindrom615/sdkman/sdk"
	"strings"
)

type envVar struct {
	name string
	val  string
}

func exportBash(paths []string, envVars []envVar) (script string) {
	res := fmt.Sprintf("export PATH=%s:$PATH\n", strings.Join(paths, ":"))
	for _, v := range envVars {
		res += fmt.Sprintf("export %s=%s\n", v.name, v.val)
	}
	return res
}

func exportFish(paths []string, envVars []envVar) (script string) {
	res := fmt.Sprintf("set -x PATH %s $PATH\n", strings.Join(paths, " "))
	for _, v := range envVars {
		res += fmt.Sprintf("set -x %s %s\n", v.name, v.val)
	}
	return res
}

func exportPosh(paths []string, envVars []envVar) (script string) {
	res := fmt.Sprintf("$env:Path = \"%s;\" + $env:Path;", strings.Join(paths, ";"))
	for _, v := range envVars {
		res += fmt.Sprintf("$env:%s = \"%s\";", v.name, v.val)
	}
	return res
}

func exportWindows(paths []string, envVars []envVar, sdkHome string) (script string) {
	currentPaths := []string{}
	for _, path := range getWindowsUserPath() {
		if !strings.Contains(path, sdkHome) && path != "" {
			currentPaths = append(currentPaths, path)
		}
	}

	paths = append(currentPaths, paths...)
	res := poshSetWindowsUserPath(paths)
	for _, v := range envVars {
		res += fmt.Sprintf("[Environment]::SetEnvironmentVariable(\"%s\", \"%s\", [System.EnvironmentVariableTarget]::User);", v.name, v.val)
	}
	return res
}

func RunExport(shell string, sdks []sdk.Sdk, sdkHome string) {
	paths, homes := getPathsHomes(sdks)

	if shell == "bash" || shell == "zsh" {
		fmt.Println(exportBash(paths, homes))
	} else if shell == "fish" {
		fmt.Println(exportFish(paths, homes))
	} else if shell == "powershell" || shell == "posh" {
		fmt.Println(exportPosh(paths, homes))
	} else if shell == "windows" || shell == "window" {
		fmt.Println(exportWindows(paths, homes, sdkHome))
	}
}
