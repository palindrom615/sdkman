package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/api"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/store"
	"github.com/spf13/cobra"
	"strings"
)

// Update updates available candidates
func update(cmd *cobra.Command, args []string) error {
	{
		freshCsv, netErr := api.GetAll(registry)
		if netErr != nil {
			return errors.ErrNotOnline
		}
		fresh := pkgs.NewStrSet(freshCsv...)
		cachedCsv := store.GetCandidates(sdkHome)
		cached := pkgs.NewStrSet(cachedCsv...)

		added := fresh.Difference(cached)
		obsoleted := cached.Difference(fresh)

		if added.Size() != 0 {
			fmt.Println("Adding new candidates: " + strings.Join(added.List(), ", "))
		}
		if obsoleted.Size() != 0 {
			fmt.Println("Removing obsolete candidates: " + strings.Join(obsoleted.List(), ", "))
		}
		return store.SetCandidates(sdkHome, freshCsv)
	}
}

var updateCmd = &cobra.Command{
	Use:  "update",
	RunE: update,
}
