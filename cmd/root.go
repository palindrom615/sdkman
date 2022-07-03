package cmd

import (
	"fmt"
	"github.com/palindrom615/sdkman/pkgs"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	registry string
	sdkHome  string
	insecure bool
)

var rootCmd = &cobra.Command{
	Use:          "sdkman",
	Short:        "manage various versions of SDKs",
	Version:      version,
	SilenceUsage: true,
}

func Execute() {
	home, _ := os.UserHomeDir()

	rootCmd.PersistentFlags().StringVarP(&registry, "registry", "r", "https://api.sdkman.io/2", "sdkman registry")
	rootCmd.PersistentFlags().StringVarP(&sdkHome, "sdkHome", "d", path.Join(home, ".sdkman"), "sdk install directory")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "whether allow insecure request")
	rootCmd.AddCommand(listCmd, currentCmd, updateCmd, installCmd, useCmd, exportCmd)

	pkgs.MkdirIfNotExist(sdkHome)
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
