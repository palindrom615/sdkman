package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/script"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
	"runtime"
)

// export prints lines of shell scripts setting up PATH and other environment variables like JAVA_HOME
func export(cmd *cobra.Command, args []string) error {
	sdks := sdk.CurrentSdks(sdkHome)
	if len(sdks) == 0 {
		fmt.Println("")
		return nil
	}

	shell := ""
	if len(args) == 0 {
		shell = getDefaultShell()
	} else {
		shell = args[0]
	}

	script.RunExport(shell, sdks)
	return nil
}

func getDefaultShell() string {
	if runtime.GOOS == "windows" {
		return "windows"
	} else {
		return "bash"
	}
}

var exportCmd = &cobra.Command{
	Use:  "export [shell]",
	RunE: export,
}
