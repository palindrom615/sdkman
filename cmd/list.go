package cmd

import (
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) error {
	reg := registry
	root := directory

	if len(args) == 0 {
		list, err := pkgs.GetList(reg)
		if err == nil {
			pkgs.Pager(list)
		}
		return err
	}
	candidate := args[0]
	if err := pkgs.CheckValidCand(root, candidate); err != nil {
		return err
	}
	ins := pkgs.InstalledSdks(root, candidate)
	curr, _ := pkgs.CurrentSdk(root, candidate)
	list, err := pkgs.GetVersionsList(reg, curr, ins)
	pkgs.Pager(list)
	return err
}

var listCmd = &cobra.Command{
	Use:                        "list [candidate]",
	Aliases:                    []string{"l", "ls"},
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
	RunE:                       list,
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
