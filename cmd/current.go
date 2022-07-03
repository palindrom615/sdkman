package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// Current print currently used packages
func current(cmd *cobra.Command, args []string) error {
	root := sdkHome
	if len(args) == 0 {
		sdks := sdk.CurrentSdks(root)
		if len(sdks) == 0 {
			return errors.ErrNoCurrCands
		}
		for _, sdk := range sdks {
			fmt.Printf("%s@%s\n", sdk.Candidate, sdk.Version)
		}
	} else {
		candidate := args[0]
		sdk, err := sdk.CurrentSdk(root, candidate)
		if err == nil {
			fmt.Println(sdk.Candidate + "@" + sdk.Version)
		} else {
			return errors.ErrNoCurrSdk(candidate)
		}
	}
	return nil
}

var currentCmd = &cobra.Command{
	Use:     "current [candidate]",
	Aliases: []string{"c"},
	RunE:    current,
}
