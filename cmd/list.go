package cmd

import (
	"github.com/palindrom615/sdkman/api"
	"github.com/palindrom615/sdkman/custom_errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

func list(cmd *cobra.Command, args []string) error {
	store.Update(registry)
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
	if !store.HasCandidate(candidate) {
		return custom_errors.ErrNoCand
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
