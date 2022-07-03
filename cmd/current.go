package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// Current print currently used packages
func current(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return currentAll()
	}
	return currentCandidate(args[0])
}

func currentAll() error {
	sdks := sdk.CurrentSdks(sdkHome)
	if len(sdks) == 0 {
		return errors.ErrNoCurrCands
	}
	for _, s := range sdks {
		fmt.Printf("%s@%s\n", s.Candidate, s.Version)
	}
	return nil
}

func currentCandidate(candidate string) error {
	currentSdk, err := sdk.CurrentSdk(sdkHome, candidate)
	if err != nil {
		return errors.ErrNoCurrSdk(candidate)

	}
	fmt.Println(currentSdk.Candidate + "@" + currentSdk.Version)
	return nil
}

var currentCmd = &cobra.Command{
	Use:     "current [candidate]",
	Aliases: []string{"c"},
	RunE:    current,
}
