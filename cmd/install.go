package cmd

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/sdk"
	"github.com/spf13/cobra"
)

// install package
func install(c *cobra.Command, args []string) error {
	_ = updateCmd.RunE(c, args)

	target, err := sdk.GetFromVersionString(registry, sdkHome, args[0])
	if err != nil {
		return err
	}

	if err := pkgs.CheckValidCand(sdkHome, target.Candidate); err != nil {
		return err
	}
	if target.Version == "" {
		defaultSdk, err := sdk.DefaultSdk(registry, sdkHome, target.Candidate)
		if err != nil {
			return err
		}
		target = defaultSdk
	}

	if target.IsInstalled(sdkHome) {
		return errors.ErrVerExists
	}
	if err := target.CheckValidVer(registry, sdkHome); err != nil {
		return err
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go target.Unarchive(sdkHome, archiveReady, installReady)
	if target.IsArchived(sdkHome) {
		archiveReady <- true
	} else {
		s, t, err := pkgs.GetDownload(registry, target)
		if err != nil {
			archiveReady <- false
			return err
		}
		archive := pkgs.Archive{target, t}
		go archive.Save(s, sdkHome, archiveReady)
	}
	if <-installReady == false {
		return errors.ErrVerInsFail
	}
	return target.Use(sdkHome)
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
