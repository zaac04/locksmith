package cmd

import (
	"context"
	"fmt"
	"locksmith/utilities"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var amend = &cobra.Command{
	Use:   "amend",
	Short: "make amendments to cipher file",
	Run: func(cmd *cobra.Command, args []string) {
		ctxVal := utilities.ReadCtx(cmd.Context(), CtxKey)
		cipherFile, _ := cmd.Flags().GetString("cipher")
		key, _ := cmd.Flags().GetString("key")

		if cipherFile == "" && key == "" {
			matches, err := utilities.DetectFile()
			if err != nil {
				utilities.LogIfError(err)
				return
			}

			if len(matches) == 0 {
				utilities.LogIfError(fmt.Errorf("%s", "no files found!"))
				return
			}

			start := time.Now()
			cipher, err := selectFile(matches)

			if err != nil {
				utilities.LogIfError(err)
				return
			}

			fmt.Print("Encryption Key: ")
			fmt.Scanf("%s", &key)
			ctxVal.UserTime += time.Since(start)
			cipherFile = cipher
		}

		//		_, cipherBytes, err := crypter.GetDecryptedValue(&cipherFile, &key)

		// if err != nil {
		// 	utilities.LogIfError(err)
		// 	return
		// }

		start := time.Now()
		//ui.EditFile(&cipherBytes, cipherFile, key)
		ctxVal.UserTime += time.Since(start)

		cmd.Parent().SetContext(context.WithValue(cmd.Context(), CtxKey, ctxVal))
	},
}

func init() {
	rootCmd.AddCommand(amend)
	amend.Flags().StringP("original", "o", "", "original file")
	amend.Flags().StringP("cipher", "c", "", "cipher file")
	amend.Flags().StringP("key", "k", "", "key")
}

func selectFile(matches []string) (result string, err error) {
	prompt := promptui.Select{
		Label:    "Select file to amend",
		Items:    matches,
		HideHelp: true,
	}
	_, result, err = prompt.Run()

	if err != nil {
		return
	}
	return result, err
}
