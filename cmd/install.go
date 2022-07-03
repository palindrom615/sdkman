package cmd

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// install package
func install(c *cobra.Command, args []string) error {
	err := store.Update(registry)
	if err != nil {
		return err
	}

	target, err := sdk.GetFromVersionString(registry, sdkHome, args[0])
	if err != nil {
		return err
	}

	if !store.HasCandidate(target.Candidate) {
		return errors.ErrNoCand
	}
	if target.Version == "" {
		defaultSdk, err := sdk.DefaultSdk(registry, sdkHome, target.Candidate)
		if err != nil {
			return err
		}
		target = defaultSdk
	}

	err = target.Install(registry)
	if err != nil {
		return err
	}
	return target.Use()
}

var installCmd = &cobra.Command{
	Use:     "install candidate[@version]",
	Aliases: []string{"i"},
	RunE:    install,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.ErrNoCand
		}
		return nil
	},
}
