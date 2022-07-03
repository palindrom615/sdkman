package cmd

import (
	"github.com/palindrom615/sdkman/api"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/palindrom615/sdkman/validate"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return listAll()
	}
	return listCandidate(args[0])
}

func listAll() error {
	list, err := api.GetList(registry)
	if err == nil {
		pkgs.Pager(list)
	}
	return err
}

func listCandidate(candidate string) error {
	if err := validate.CheckValidCand(sdkHome, candidate); err != nil {
		return err
	}
	installedSdk := sdk.InstalledSdks(sdkHome, candidate)
	installedVersion := make([]string, len(installedSdk))
	for _, s := range installedSdk {
		installedVersion = append(installedVersion, s.Version)
	}
	curr, _ := sdk.CurrentSdk(sdkHome, candidate)
	list, err := api.GetVersionsList(registry, curr.Candidate, curr.Version, installedVersion)
	pkgs.Pager(list)
	return err
}

var listCmd = &cobra.Command{
	Use:     "list [candidate]",
	Aliases: []string{"l", "ls"},
	RunE:    list,
}
