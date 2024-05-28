package cmd

import (
	"context"
	"fmt"
	ui "locksmith/Ui"
	"locksmith/crypter"
	"locksmith/file"
	"locksmith/utilities"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "encrypt the file specifed",
	Run: func(cmd *cobra.Command, args []string) {

		ctxVal := utilities.ReadCtx(cmd.Context(), CtxKey)
		stage, _ := cmd.Flags().GetString("stage")
		key, _ := cmd.Flags().GetString("key")

		if stage != "" {
			stage = utilities.GetCipherName(stage)
		}

		Exist, err := utilities.DoFileExist(stage)

		if err != nil {
			utilities.LogIfError(err)
			return
		}

		if Exist && key == "" {
			fmt.Print("Encryption Key: ")
			fmt.Scanf("%s", &key)
		}

		if strings.Contains(key, ".locksmith.key") {
			data, err := file.ReadFile(key)
			if err != nil {
				utilities.LogIfError(err)
				return
			}
			key = string(data)
		} else if key == "" && !Exist {
			start := time.Now()
			encryption := selectEncryption()
			ctxVal.UserTime += time.Since(start)

			lock, err := crypter.New(encryption)
			if err != nil {
				utilities.LogIfError(err)
				return
			}
			key = lock.GetKey()
			defer fmt.Println("Key:", key)
			defer file.WriteFile(stage+".key", []byte(key))
		}

		cypherByte := []byte{}

		if Exist {
			_, cypherByte, err = crypter.GetDecryptedValue(&stage, &key)
			if err != nil {
				utilities.LogIfError(err)
				return
			}
		}

		start := time.Now()
		ui.EditFile(&cypherByte, stage, key)
		ctxVal.UserTime += time.Since(start)
		cmd.Parent().SetContext(context.WithValue(cmd.Context(), CtxKey, ctxVal))

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("stage", "s", "", "name of file to encrypt")
	addCmd.Flags().StringP("key", "k", "", "key to encrypt")
	addCmd.MarkFlagRequired("stage")
}
