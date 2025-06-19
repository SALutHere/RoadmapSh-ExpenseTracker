package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense tracker",
	Long:  "Program to track expenses. Supports listing, adding, updating and deleting expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to expense-tracker!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(summaryCmd)
}
