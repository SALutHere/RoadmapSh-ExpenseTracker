package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		expense := expenses.NewExpense(description, amount, category)
		if err := expenses.AddExpense(expense); err != nil {
			log.Fatalf("error when adding a new expense: %v", err)
		}
		fmt.Printf("A new expense was added successfully! ID: %s", expense.ID.String())
	},
}

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Expense description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Expense amount")
	addCmd.Flags().StringVarP(&category, "category", "c", "", "Expense category")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("category")
}
