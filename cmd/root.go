package cmd

import (
	"fmt"
	"locksmith/version"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var VersionFlag bool

var rootCmd = &cobra.Command{
	Use:   "locksmith",
	Short: "locksmith helps you manage envs easily!",
}

func Execute() {

	start := time.Now()
	print("\n")
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "\nOperation Completed in %s\n", time.Since(start))
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = version.PrintVersion()
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true
}
