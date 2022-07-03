package cmd

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

// use make symbolic link named "current" to installed package.
func use(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.ErrNoCand
	}
	sdk, err := pkgs.Arg2sdk(registry, sdkHome, args[0])
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
}
