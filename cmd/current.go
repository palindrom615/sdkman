package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

// Current print currently used packages
func current(cmd *cobra.Command, args []string) error {
	root := directory
	if len(args) == 0 {
		sdks := pkgs.CurrentSdks(root)
		if len(sdks) == 0 {
			return errors.ErrNoCurrCands
		}
		for _, sdk := range sdks {
			fmt.Printf("%s@%s\n", sdk.Candidate, sdk.Version)
		}
	} else {
		candidate := args[0]
		sdk, err := pkgs.CurrentSdk(root, candidate)
		if err == nil {
			fmt.Println(sdk.Candidate + "@" + sdk.Version)
		} else {
			return errors.ErrNoCurrSdk(candidate)
		}
	}
	return nil
}

var currentCmd = &cobra.Command{
	Use:                        "current [candidate]",
	Aliases:                    []string{"c"},
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
	RunE:                       current,
	Run:                        nil,
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
