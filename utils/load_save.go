package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadsPasswords() map[string]string {
	passwords := make(map[string]string)
	file, err := os.ReadFile("passwords.json")
	if err != nil {
		if os.IsNotExist(err) {
			return passwords
		}
		fmt.Println("Error Reading file:", err)
		return nil
	}
	err = json.Unmarshal(file, &passwords)
	if err != nil {
		fmt.Println("Error unmarshling JSON", err)
		return nil
	}
	// fmt.Println(passwords)
	return passwords
}
