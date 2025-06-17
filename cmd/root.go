package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense tracker",
	Long:  "Program to track expenses. Supports adding and viewing expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to expense-tracker!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
