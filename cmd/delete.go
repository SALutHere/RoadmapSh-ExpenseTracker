package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deleteId string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing expense",
	Run: func(cmd *cobra.Command, args []string) {
		if err := expenses.DeleteExpense(uuid.MustParse(deleteId)); err != nil {
			log.Fatalf("error when deleting a new expense: %v\n", err)
		}
		fmt.Printf("Expense was deleted successfully! ID: %s\n", deleteId)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteId, "id", "i", "", "Expense ID")

	deleteCmd.MarkFlagRequired("id")
}
