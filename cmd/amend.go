package cmd

import (
	"context"
	"fmt"
	"locksmith/crypter"
	"locksmith/utilities"
	"path/filepath"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var amend = &cobra.Command{
	Use:   "amend",
	Short: "make amendments to cipher file",
	Run: func(cmd *cobra.Command, args []string) {
		ctxVal := utilities.ReadCtx(cmd.Context(), CtxKey)
		original, _ := cmd.Flags().GetString("original")
		cipherFile, _ := cmd.Flags().GetString("cipher")
		key, _ := cmd.Flags().GetString("key")

		if original == "" && cipherFile == "" && key == "" {
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
			cipher, file, err := selectFile(matches)

			if err != nil {
				utilities.LogIfError(err)
				return
			}

			fmt.Print("Encryption Key: ")
			fmt.Scanf("%s", &key)
			ctxVal.UserTime += time.Since(start)
			original = file
			cipherFile = cipher
		}
		lock, change, err := crypter.FinDiff(original, cipherFile, key)

		if err != nil {
			utilities.LogIfError(err)
			return
		}

		if change {
			start := time.Now()
			confirm, err := confirmAmend()
			if err != nil {
				utilities.LogIfError(err)
				return
			}
			ctxVal.UserTime += time.Since(start)
			if confirm {
				crypter.Encrypt(lock, original)
			}
		} else {
			fmt.Println("No Change Detected!")
		}
		cmd.Parent().SetContext(context.WithValue(cmd.Context(), CtxKey, ctxVal))
	},
}

func init() {
	rootCmd.AddCommand(amend)
	amend.Flags().StringP("original", "o", "", "original file")
	amend.Flags().StringP("cipher", "c", "", "cipher file")
	amend.Flags().StringP("key", "k", "", "key")
}

func selectFile(matches []string) (result string, file string, err error) {
	prompt := promptui.Select{
		Label:    "Select file to amend",
		Items:    matches,
		HideHelp: true,
	}
	_, result, err = prompt.Run()

	if err != nil {
		return
	}
	ext := filepath.Ext(result)
	file, _ = strings.CutSuffix(result, ext)
	return result, file, err
}

func confirmAmend() (proceed bool, err error) {
	options := []string{"Yes", "No"}

	mapOptions := map[string]bool{
		"Yes": true,
		"No":  false,
	}
	prompt := promptui.Select{
		Label:    "Confirm Amend",
		Items:    options,
		HideHelp: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return
	}

	if val, ok := mapOptions[result]; ok {
		proceed = val
	}
	return
}
