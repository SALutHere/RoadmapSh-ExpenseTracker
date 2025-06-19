package expenses

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/google/uuid"
)

// Returns table format of given expenses
func GetExpensesTable(expenses map[uuid.UUID]Expense) string {
	var buffer bytes.Buffer
	var writer = tabwriter.NewWriter(&buffer, 0, 8, 1, '\t', 0)

	fmt.Fprintln(writer, "ID\tDescription\tAmount\tCategory\tDate")
	for _, expense := range expenses {
		fmt.Fprintf(writer, "%s\t%s\t$%.2f\t%s\t%s\n",
			expense.ID.String(),
			expense.Description,
			expense.Amount,
			expense.Category,
			expense.Date.Format("02-01-2006 15:04:05"),
		)
	}
	writer.Flush()

	return buffer.String()
}

// Returns expenses filtered by day
func FilterExpensesByDay(expenses map[uuid.UUID]Expense, day int) map[uuid.UUID]Expense {
	for id, expense := range expenses {
		if int(expense.Date.Day()) != day {
			delete(expenses, id)
		}
	}
	return expenses
}

// Returns expenses filtered by month
func FilterExpensesByMonth(expenses map[uuid.UUID]Expense, month int) map[uuid.UUID]Expense {
	for id, expense := range expenses {
		if int(expense.Date.Month()) != month {
			delete(expenses, id)
		}
	}
	return expenses
}

// Returns expenses filtered by year
func FilterExpensesByYear(expenses map[uuid.UUID]Expense, year int) map[uuid.UUID]Expense {
	for id, expense := range expenses {
		if int(expense.Date.Year()) != year {
			delete(expenses, id)
		}
	}
	return expenses
}

// Returns expenses filtered by category
func FilterExpensesByCategory(expenses map[uuid.UUID]Expense, category string) map[uuid.UUID]Expense {
	for id, expense := range expenses {
		if expense.Category != category {
			delete(expenses, id)
		}
	}
	return expenses
}

// Returns total expenses of given expenses
func GetTotalExpenses(expenses map[uuid.UUID]Expense) float64 {
	var total float64

	for _, expense := range expenses {
		total += expense.Amount
	}

	return total
}
