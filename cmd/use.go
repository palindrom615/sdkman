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
	sdk, err := pkgs.Arg2sdk(registry, directory, args[0])
	if err != nil {
		return err
	}
	if !sdk.IsInstalled(directory) {
		return errors.ErrVerNotIns
	}
	return sdk.Use(directory)
}

var useCmd = &cobra.Command{
	Use:                        "use candidate@version",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "",
	Long:                       "",
	Example:                    "",
	ValidArgs:                  nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Hidden:                     false,
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run:                        nil,
	RunE:                       use,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}
