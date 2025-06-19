package cmd

import (
	"expense-tracker/internal/expenses"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var summaryDay int
var summaryMonth int
var summaryYear int
var summaryCategory string

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get sum of expenses by specified filters",
	Run: func(cmd *cobra.Command, args []string) {
		exps, err := expenses.GetExpensesFromFile()
		if err != nil {
			log.Fatalf("error when getting sum of expenses: %v", err)
		}

		if summaryDay != 0 {
			exps = expenses.FilterExpensesByDay(exps, summaryDay)
		}
		if summaryMonth != 0 {
			exps = expenses.FilterExpensesByMonth(exps, summaryMonth)
		}
		if summaryYear != 0 {
			exps = expenses.FilterExpensesByYear(exps, summaryYear)
		}
		if summaryCategory != "" {
			exps = expenses.FilterExpensesByCategory(exps, summaryCategory)
		}

		if len(exps) != 0 {
			total := expenses.GetTotalExpenses(exps)
			fmt.Printf("Total expenses: $%.2f", total)
		} else {
			fmt.Println("Suitable results were not found")
		}

	},
}

func init() {
	summaryCmd.Flags().IntVarP(&summaryDay, "day", "d", 0, "Expenses filtering day")
	summaryCmd.Flags().IntVarP(&summaryMonth, "month", "m", 0, "Expenses filtering month")
	summaryCmd.Flags().IntVarP(&summaryYear, "year", "y", 0, "Expenses filtering year")
	summaryCmd.Flags().StringVarP(&summaryCategory, "category", "c", "", "Expenses filtering category")
}
