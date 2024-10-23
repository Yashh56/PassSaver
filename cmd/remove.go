package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var key string

var removecmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove an stored username and password",
	Run: func(cmd *cobra.Command, args []string) {
		removeUsernameAndPassword(key)
	},
}

func init() {
	rootCmd.AddCommand(removecmd)
	removecmd.Flags().StringVarP(&key, "key", "u", "", "key for the account")
	removecmd.MarkFlagRequired("key")
}

func removeUsernameAndPassword(key string) {

	username := make(map[string]string)

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

	if _, exists := username[key]; exists {
		delete(username, key)
		fmt.Println("Key value has been deleted!")
		updatedData, err := json.MarshalIndent(username, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		err = ioutil.WriteFile("passwords.json", updatedData, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	} else {
		fmt.Println("No data found!")
	}
}
