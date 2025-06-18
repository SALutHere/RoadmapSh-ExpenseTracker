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
		fmt.Fprintf(writer, "%s\t%s\t%.2f\t%s\t%s\n",
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
