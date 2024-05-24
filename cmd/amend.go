package cmd

import (
	"locksmith/crypter"
	"locksmith/utilities"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var amend = &cobra.Command{
	Use:   "amend",
	Short: "make amendments to cipher file",
	Run: func(cmd *cobra.Command, args []string) {
		matches, err := utilities.DetectFile()

		if err != nil {
			utilities.LogIfError(err)
			return
		}

		original, _ := cmd.Flags().GetString("original")
		cipherFile, _ := cmd.Flags().GetString("cipher")
		key, _ := cmd.Flags().GetString("key")

		if original == "" && cipherFile == "" && key != "" {

			cipher, file := selectFile(matches)

			original = file
			cipherFile = cipher

		}
		crypter.Amend(original, cipherFile, key)

	},
}

func init() {
	rootCmd.AddCommand(amend)
	amend.Flags().StringP("original", "o", "", "original file")
	amend.Flags().StringP("cipher", "c", "", "cipher file")
	amend.Flags().StringP("key", "k", "", "key")
	amend.MarkFlagRequired("key")
}

func selectFile(matches []string) (string, string) {
	prompt := promptui.Select{
		Label:    "Select an encryption to use:",
		Items:    matches,
		HideHelp: true,
	}
	_, result, _ := prompt.Run()

	ext := filepath.Ext(result)
	file, _ := strings.CutSuffix(result, ext)
	return result, file
}
