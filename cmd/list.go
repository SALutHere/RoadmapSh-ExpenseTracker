package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the expenses",
	Run: func(cmd *cobra.Command, args []string) {
		exps, err := expenses.GetExpensesFromFile()
		if err != nil {
			log.Fatalf("error when listing expenses: %v\n", err)
		}
		output := expenses.GetExpensesTable(exps)
		fmt.Println(output)
	},
}
