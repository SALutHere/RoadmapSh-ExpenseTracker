package expenses

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID
	Date        time.Time
	Description string
	Amount      float64
	Category    string
}

// Returns a new expense by specified description, amount and category
func NewExpense(descriprtion string, amount float64, category string) Expense {
	id := uuid.New()
	date := time.Now()
	expense := Expense{
		ID:          id,
		Date:        date,
		Description: descriprtion,
		Amount:      amount,
		Category:    category,
	}
	return expense
}
