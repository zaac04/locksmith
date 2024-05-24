package cmd

import (
	"locksmith/crypter"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypts the file passed in argument",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		key, _ := cmd.Flags().GetString("key")
		crypter.DecryptFile(filename, key)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringP("key", "k", "", "key used to decrypt")
	decryptCmd.Flags().StringP("file", "f", "", "filename")

	decryptCmd.MarkFlagRequired("file")
	decryptCmd.MarkFlagsRequiredTogether("key", "file")
}
