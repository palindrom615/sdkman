package cmd

import (
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) error {
	reg := registry
	root := sdkHome

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
	Use:     "list [candidate]",
	Aliases: []string{"l", "ls"},
	RunE:    list,
}
