package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "password-manager",
	Short: "A simple CLI password manager",
	Long:  `A CLI password manager written in Go using Cobra for storing passwords securely.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
