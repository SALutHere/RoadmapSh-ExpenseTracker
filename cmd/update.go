package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var updateId string
var updateDescription string
var updateAmount float64
var updateCategory string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		expense := expenses.Expense{
			ID:          uuid.MustParse(updateId),
			Description: updateDescription,
			Amount:      updateAmount,
			Category:    updateCategory,
		}
		if err := expenses.UpdateExpense(expense); err != nil {
			log.Fatalf("error when updating an expense: %v", err)
		}
		fmt.Printf("An expense was updated successfully! ID: %v", expense.ID)
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateId, "id", "i", "", "Expense ID")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "Expense description")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0, "Expense amount")
	updateCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "Expense category")

	updateCmd.MarkFlagRequired("id")
}
