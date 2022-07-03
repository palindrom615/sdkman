package cmd

import (
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return listAll()
	}
	return listCandidate(args[0])
}

func listAll() error {
	list, err := pkgs.GetList(registry)
	if err == nil {
		pkgs.Pager(list)
	}
	return err
}

func listCandidate(candidate string) error {
	if err := pkgs.CheckValidCand(sdkHome, candidate); err != nil {
		return err
	}
	ins := sdk.InstalledSdks(sdkHome, candidate)
	curr, _ := sdk.CurrentSdk(sdkHome, candidate)
	list, err := pkgs.GetVersionsList(registry, curr, ins)
	pkgs.Pager(list)
	return err
}

var listCmd = &cobra.Command{
	Use:     "list [candidate]",
	Aliases: []string{"l", "ls"},
	RunE:    list,
}
