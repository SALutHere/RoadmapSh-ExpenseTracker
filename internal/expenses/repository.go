package expenses

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

// Path of the storage file of expenses
const expensesFilePath = "storage/expenses.json"

// Returns the current expenses map taken from the storage file
func GetExpensesFromFile() (map[uuid.UUID]Expense, error) {
	var expensesFile *os.File

	if _, err := os.Stat(expensesFilePath); errors.Is(err, os.ErrNotExist) {
		expensesFile, err = os.Create(expensesFilePath)
		if err != nil {
			return nil, fmt.Errorf("error creating expenses file: %v", err)
		}
	} else {
		expensesFile, err = os.Open(expensesFilePath)
		if err != nil {
			return nil, fmt.Errorf("error opening expenses file for reading: %v", err)
		}
	}
	defer expensesFile.Close()

	expenses := make(map[uuid.UUID]Expense)
	decoder := json.NewDecoder(expensesFile)
	if err := decoder.Decode(&expenses); err != nil && err.Error() != "EOF" {
		return nil, fmt.Errorf("error decoding expenses file: %v", err)
	}
	return expenses, nil
}

// Writes given expenses map to the storage file
func WriteExpensesToFile(expenses map[uuid.UUID]Expense) error {
	expensesFile, err := os.OpenFile(expensesFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("error opening expenses file for writing: %v", err)
	}
	defer expensesFile.Close()

	encoder := json.NewEncoder(expensesFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&expenses); err != nil {
		return fmt.Errorf("error encoding expenses file: %v", err)
	}

	return nil
}

// Saves given expense to the storage file
func AddExpense(expense Expense) error {
	if expense.Amount < 0 {
		log.Fatal("negative amount is not valid")
	}

	expenses, err := GetExpensesFromFile()
	if err != nil {
		return err
	}

	expenses[expense.ID] = expense

	if err := WriteExpensesToFile(expenses); err != nil {
		return fmt.Errorf("error adding an expense to file: %v", err)
	}

	return nil
}

// Updates existing expense in the storage file
func UpdateExpense(expense Expense) error {
	if expense.Amount < 0 {
		log.Fatal("negative amount is not valid")
	}

	expenses, err := GetExpensesFromFile()
	if err != nil {
		return err
	}

	if exp, ok := expenses[expense.ID]; ok {
		if expense.Description != "" {
			exp.Description = expense.Description
		}
		if expense.Amount != 0 {
			exp.Amount = expense.Amount
		}
		if expense.Category != "" {
			exp.Category = expense.Category
		}

		expenses[expense.ID] = exp
	} else {
		return fmt.Errorf("expense with ID %s not found", expense.ID)
	}

	if err := WriteExpensesToFile(expenses); err != nil {
		return fmt.Errorf("error adding an expense to file: %v", err)
	}

	return nil
}

// Deletes existing expense from the storage file by its ID
func DeleteExpense(id uuid.UUID) error {
	expenses, err := GetExpensesFromFile()
	if err != nil {
		return err
	}

	if _, ok := expenses[id]; ok {
		delete(expenses, id)
	} else {
		return fmt.Errorf("expense with ID %s not found", id.String())
	}

	if err := WriteExpensesToFile(expenses); err != nil {
		return fmt.Errorf("error adding an expense to file: %v", err)
	}

	return nil
}
