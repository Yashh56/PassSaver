package cmd

import (
	"fmt"

	"github.com/Yashh56/passSaver/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored usernames and passwords",
	Run: func(cmd *cobra.Command, args []string) {
		listPasswords()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listPasswords() {
	encryptionKey := "thisisasecretkey"
	passwords := utils.LoadsPasswords()

	fmt.Println("Stored passwords:")
	for username, encryptedPassword := range passwords {
		decryptedPassword, err := utils.Decrypt(encryptionKey, encryptedPassword)
		if err != nil {
			fmt.Println("Error decrypting password for", username)
			continue
		}
		fmt.Printf("Username: %s, Password: %s\n", username, decryptedPassword)
	}
}
