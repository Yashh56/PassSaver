package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yashh56/passSaver/utils"
	"github.com/spf13/cobra"
)

var username, password string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new username and password",
	Run: func(cmd *cobra.Command, args []string) {
		addPassword(username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the account")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "Password for the account")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("password")
}

func addPassword(username, password string) {
	passwords := utils.LoadsPasswords()
	encryptionKey := "thisisasecretkey"
	encryptionPassword, err := utils.Encrypt(encryptionKey, password)
	if err != nil {
		fmt.Println("Error encrypting password", err)
		return
	}
	passwords[username] = encryptionPassword
	file, err := os.Create("passwords.json")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(passwords, "", " ")
	if err != nil {
		fmt.Println("Error marshling error", err)
		return
	}
	file.Write(jsonData)
	fmt.Println("Password added successfully!")
}
