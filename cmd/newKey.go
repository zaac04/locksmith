package cmd

import (
	"context"
	"fmt"
	"locksmith/crypter"
	"locksmith/utilities"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var genKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate Secret key for encryption",
	Run: func(cmd *cobra.Command, args []string) {

		ctxVal := utilities.ReadCtx(cmd.Context(), CtxKey)

		//user start time
		start := time.Now()
		encryption := selectOption()
		ctxVal.UserTime = time.Since(start)

		cmd.Parent().SetContext(context.WithValue(cmd.Context(), CtxKey, ctxVal))
		Lock, err := crypter.New(encryption)
		utilities.LogIfError(err)

		fmt.Println("key:", Lock.GetKey())
	},
}

func init() {
	rootCmd.AddCommand(genKeyCmd)
}

func selectOption() (encryption int) {
	options := []string{"AES 128bit", "AES 192bit", "AES 256bit (default)"}

	mapOptions := map[string]int{
		"AES 128bit":           128,
		"AES 192bit":           192,
		"AES 256bit (default)": 256,
	}

	prompt := promptui.Select{
		Label:    "\nConfirm Amend:",
		Items:    options,
		HideHelp: true,
	}

	_, result, _ := prompt.Run()

	if val, ok := mapOptions[result]; ok {
		return val
	}

	return 256
}
