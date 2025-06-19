package cmd

import (
	"expense-tracker/internal/expenses"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears your expenses list",
	Run: func(cmd *cobra.Command, args []string) {
		emptiness := map[uuid.UUID]expenses.Expense{}
		expenses.WriteExpensesToFile(emptiness)
	},
}
