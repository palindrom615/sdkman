package cmd

import (
	"fmt"
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
	registry  string
	directory string
	insecure  bool
)

var rootCmd = &cobra.Command{
	Use:     "sdkman",
	Short:   "manage various versions of SDKs",
	Version: version,
}

func Execute() {
	home, _ := os.UserHomeDir()

	rootCmd.PersistentFlags().StringVarP(&registry, "registry", "r", "https://api.sdkman.io/2", "sdkman registry")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", path.Join(home, ".sdkman"), "sdkman registry")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "whether allow insecure request")
	rootCmd.AddCommand(listCmd, currentCmd, updateCmd, installCmd, useCmd, exportCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
