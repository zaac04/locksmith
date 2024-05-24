package cmd

import (
	"locksmith/crypter"

	"github.com/spf13/cobra"
)

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "checks whether the file and cipher hash matches",
	Run: func(cmd *cobra.Command, args []string) {
		original, _ := cmd.Flags().GetString("original")
		cipherFile, _ := cmd.Flags().GetString("cipher")
		crypter.MatchCheckSum(original, cipherFile)
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringP("original", "o", "", "name of original file")
	checksumCmd.Flags().StringP("cipher", "c", "", "name of cipher file")
	checksumCmd.MarkFlagRequired("original")
	checksumCmd.MarkFlagsRequiredTogether("original", "cipher")
}
