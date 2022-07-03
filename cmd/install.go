package cmd

import (
	"errors"
	"github.com/palindrom615/sdkman/custom_errors"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
	"sync"
)

// install package
func install(c *cobra.Command, args []string) error {
	err := store.Update(registry)
	if err != nil {
		return err
	}

	targets := []sdk.Sdk{}
	errMsg := ""
	for _, arg := range args {
		target, e := sdk.GetFromVersionString(registry, sdkHome, arg)
		targets = append(targets, target)
		if e != nil {
			errMsg = errMsg + arg + ": " + e.Error() + "\n"
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}

	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		go func(target sdk.Sdk) {
			defer wg.Done()
			if target.Version == "" {
				defaultSdk, err := sdk.DefaultSdk(registry, sdkHome, target.Candidate)
				if err != nil {
					println(target.ToString() + ": " + err.Error())
					return
				}
				target = defaultSdk
			}

			err := target.Install(registry)
			if err != nil {
				println(target.ToString() + ": " + err.Error())
			}
			target.Use()
		}(target)
	}
	wg.Wait()
	return nil
}

var installCmd = &cobra.Command{
	Use:     "install candidate[@version]...",
	Aliases: []string{"i"},
	RunE:    install,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return custom_errors.ErrNoCand
		}
		return nil
	},
}
