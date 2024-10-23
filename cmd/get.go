package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yashh56/passSaver/utils"
	"github.com/spf13/cobra"
)

var user string

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"u"},
	Short:   "Get user and password",
	Run: func(cmd *cobra.Command, args []string) {
		getSinglePassword(user)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&user, "user", "u", "", "Username for the account")
	getCmd.MarkFlagRequired("user")
}

func getSinglePassword(user string) {

	username := make(map[string]string)
	encryptionKey := "thisisasecretkey"

	file, err := os.ReadFile("passwords.json")

	if err != nil {
		fmt.Println("Remove user error", err)
		return
	}

	err = json.Unmarshal(file, &username)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}

	if _, exists := username[user]; exists {
		decryptedPassword, err := utils.Decrypt(encryptionKey, username[user])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s: %s\n", user, decryptedPassword)
	} else {
		fmt.Println("No data found!")
	}
}
