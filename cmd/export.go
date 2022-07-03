package cmd

import (
	"fmt"
	"github.com/yargevad/filepathx"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

// export prints lines of shell scripts setting up PATH and other environment variables like JAVA_HOME
func export(cmd *cobra.Command, args []string) error {
	shell := ""
	if len(args) == 0 {
		if runtime.GOOS == "windows" {
			shell = "windows"
		} else {
			shell = "bash"
		}
	} else {
		shell = args[0]
	}
	sdks := pkgs.CurrentSdks(sdkHome)
	if len(sdks) == 0 {
		fmt.Println("")
		return nil
	}
	paths := []string{}
	homes := []envVar{}
	for _, sdk := range sdks {
		candHome := path.Join(sdkHome, "candidates", sdk.Candidate, "current")

		binPath := path.Join(candHome, "bin")
		homePath := candHome
		if sdk.Candidate == "java" && runtime.GOOS == "darwin" {
			// fix for macOS
			matches, _ := filepathx.Glob(fmt.Sprintf("%s/**/javac", homePath))
			binPath = filepath.Join(matches[0], "..")
			homePath = filepath.Join(binPath, "..")
		}
		paths = append(paths, binPath)
		homes = append(homes, envVar{fmt.Sprintf("%s_HOME", strings.ToUpper(sdk.Candidate)), homePath})
	}

	if shell == "bash" || shell == "zsh" {
		fmt.Println(exportBash(paths, homes))
	} else if shell == "fish" {
		fmt.Println(exportFish(paths, homes))
	} else if shell == "powershell" || shell == "posh" {
		fmt.Println(exportPosh(paths, homes))
	} else if shell == "windows" || shell == "window" {
		fmt.Println(exportWindows(paths, homes))
	}
	return nil
}

var exportCmd = &cobra.Command{
	Use:  "export [shell]",
	RunE: export,
}
