package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/pkgs/strset"
	"github.com/palindrom615/sdkman/store"
	"github.com/spf13/cobra"
	"strings"
)

// Update updates available candidates
func update(cmd *cobra.Command, args []string) error {
	{
		freshCsv, netErr := pkgs.GetAll(registry)
		if netErr != nil {
			return errors.ErrNotOnline
		}
		fresh := strset.New(freshCsv...)
		cachedCsv := store.GetCandidates(sdkHome)
		cached := strset.New(cachedCsv...)

		added := fresh.Difference(cached)
		obsoleted := cached.Difference(fresh)

		if added.Size() != 0 {
			fmt.Println("Adding new candidates: " + strings.Join(added.List(), ", "))
		}
		if obsoleted.Size() != 0 {
			fmt.Println("Removing obsolete candidates: " + strings.Join(obsoleted.List(), ", "))
		}
		_ = store.SetCandidates(sdkHome, freshCsv)
		return nil
	}
}

var updateCmd = &cobra.Command{
	Use:  "update",
	RunE: update,
}
