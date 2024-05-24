package cmd

import (
	"locksmith/crypter"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt the file specifed",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")
		crypter.EncryptFile(filename, key)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringP("file", "f", "", "name of file to encrypt")
	encryptCmd.Flags().StringP("key", "k", "", "key to encrypt")
	encryptCmd.MarkFlagRequired("file")
	encryptCmd.MarkFlagsRequiredTogether("key", "file")
}
