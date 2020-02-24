package cmd

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/spf13/cobra"
)

// install package
func install(c *cobra.Command, args []string) error {
	_ = updateCmd.RunE(c, args)

	if len(args) == 0 {
		return errors.ErrNoCand
	}
	target, err := pkgs.Arg2sdk(registry, directory, args[0])
	if err != nil {
		return err
	}

	pkgs.MkdirIfNotExist(directory)
	if err := pkgs.CheckValidCand(directory, target.Candidate); err != nil {
		return err
	}
	if target.Version == "" {
		defaultSdk, err := pkgs.DefaultSdk(registry, directory, target.Candidate)
		if err != nil {
			return err
		}
		target = defaultSdk
	}

	if target.IsInstalled(directory) {
		return errors.ErrVerExists
	}
	if err := target.CheckValidVer(registry, directory); err != nil {
		return err
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go target.Unarchive(directory, archiveReady, installReady)
	if target.IsArchived(directory) {
		archiveReady <- true
	} else {
		s, t, err := pkgs.GetDownload(registry, target)
		if err != nil {
			archiveReady <- false
			return err
		}
		archive := pkgs.Archive{target, t}
		go archive.Save(s, directory, archiveReady)
	}
	if <-installReady == false {
		return errors.ErrVerInsFail
	}
	return target.Use(directory)
}

var installCmd = &cobra.Command{
	Use:                        "install candidate[@version]",
	Aliases:                    []string{"i"},
	SuggestFor:                 nil,
	Short:                      "",
	Long:                       "",
	Example:                    "",
	ValidArgs:                  nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Hidden:                     false,
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run:                        nil,
	RunE:                       install,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}
