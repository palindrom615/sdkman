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
	Use:   "export [shell]",
	Short: "'sdk export [shell]' command print script putting sdk paths in system environment variables",
	Long: `'sdk export [shell]' command print script putting sdk paths in system environment variables.
	To apply script for current shell only,
	
	powershell:
		> Invoke-Expression (sdk export posh)
	bash:
		> eval $(sdk export bash)
	zsh:
		> eval $(sdk export zsh)
	fish:
		> eval (sdk export fish)


	To apply script permanently,

	windows (user-wide global):
		> Invoke-Expression (sdk export windows)
	powershell:
		> Add-Content $Profile "Invoke-Expression (sdk export posh)"
	bash:
		> echo "eval \$(sdk export bash)" >> ~/.bashrc
	zsh:
		> echo "eval \$(sdk export zsh)" >> ~/.zshrc
	fish:
		> echo "eval (sdk export fish)" >> ~/.config/config.fish	
	`,
	RunE: export,
}
