package cmd

import (
	"fmt"
	"locksmith/crypter"
	"locksmith/utilities"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var genKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		encryption := selectOption()
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
		Label:    "Select an encryption to use:",
		Items:    options,
		HideHelp: true,
	}

	_, result, _ := prompt.Run()

	if val, ok := mapOptions[result]; ok {
		return val
	}

	return 256
}
