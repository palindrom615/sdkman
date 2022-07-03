package cmd

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// use make symbolic link named "current" to installed package.
func use(cmd *cobra.Command, args []string) error {
	sdk, err := sdk.GetFromVersionString(registry, sdkHome, args[0])
	if err != nil {
		return err
	}
	if !sdk.IsInstalled(sdkHome) {
		return errors.ErrVerNotIns
	}
	return sdk.Use(sdkHome)
}

var useCmd = &cobra.Command{
	Use:  "use candidate@version",
	RunE: use,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.ErrNoCand
		}
		return nil
	},
}
