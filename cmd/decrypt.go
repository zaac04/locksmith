package cmd

import (
	"locksmith/crypter"
	"locksmith/file"
	"locksmith/utilities"
	"strings"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypts the file passed in argument",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")

		filename = utilities.GetCipherName(filename)

		key, _ := cmd.Flags().GetString("key")
		if strings.Contains(key, ".locksmith.key") {
			data, err := file.ReadFile(key)
			if err != nil {
				utilities.LogIfError(err)
				return
			}
			key = string(data)
		}
		crypter.DecryptFile(filename, key)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringP("key", "k", "", "key used to decrypt")
	decryptCmd.Flags().StringP("file", "s", "", "filename")

	decryptCmd.MarkFlagRequired("file")
	decryptCmd.MarkFlagsRequiredTogether("key", "file")
}
