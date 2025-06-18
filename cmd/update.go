package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		expense := expenses.Expense{
			ID:          uuid.MustParse(id),
			Description: description,
			Amount:      amount,
			Category:    category,
		}
		if err := expenses.UpdateExpense(expense); err != nil {
			log.Fatalf("error when updating an expense")
		}
		fmt.Printf("An expense was updated successfully! ID: %v", expense.ID)
	},
}

func init() {
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "Expense ID")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "Expense description")
	updateCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Expense amount")
	updateCmd.Flags().StringVarP(&category, "category", "c", "", "Expense category")

	updateCmd.MarkFlagRequired("id")
}
