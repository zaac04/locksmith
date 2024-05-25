package cmd

import (
	"context"
	"locksmith/utilities"
	"locksmith/version"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var VersionFlag bool
var CtxKey utilities.CtxKey = "key"

var rootCmd = &cobra.Command{
	Use:   "locksmith",
	Short: "locksmith helps you manage envs easily!",
}

func Execute() {

	var val = utilities.Val{
		Time: time.Now(),
	}

	ctx := context.WithValue(context.Background(), CtxKey, val)
	defer utilities.RecoverFromPanic()

	rootCmd.SetContext(ctx)
	err := rootCmd.Execute()

	utilities.PrintOperationTime(rootCmd.Context(), CtxKey)
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Version = version.PrintVersion()
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true
}
