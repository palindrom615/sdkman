package cmd

import (
	"github.com/palindrom615/sdkman/custom_errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// use make symbolic link named "current" to installed package.
func use(cmd *cobra.Command, args []string) error {
	store.Update(registry)
	sdk, err := sdk.GetFromVersionString(registry, sdkHome, args[0])
	if err != nil {
		return err
	}
	if !sdk.IsInstalled() {
		return custom_errors.ErrVerNotIns
	}
	return sdk.Use()
}

var useCmd = &cobra.Command{
	Use:  "use candidate@version",
	RunE: use,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return custom_errors.ErrNoCand
		}
		return nil
	},
}
