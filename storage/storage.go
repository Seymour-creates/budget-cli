package storage

import (
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/expenses"
	"os"
	"time"
)

func UpdateExpenses(newExpenses ...expenses.Expense) error {
	existingExpenses, err := LoadExpenses()
	if err != nil {
		return err
	}
	combinedExpense := append(existingExpenses, newExpenses...)
	return SaveExpenses(combinedExpense)
}

func SaveExpenses(expenses expenses.Expenses) error {
	filePath, err := _getExpenseFilePath()
	if err != nil {
		return err
	}
	return _saveExpensesToFile(filePath, expenses)
}

func LoadExpenses() (expenses.Expenses, error) {
	filePath, err := _getExpenseFilePath()
	if os.IsNotExist(err) {
		return expenses.Expenses{}, nil
	} else if err != nil {
		return nil, err
	}
	return _loadExpenseFromFile(filePath)
}

func _saveExpensesToFile(filename string, expenses expenses.Expenses) error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func _loadExpenseFromFile(filename string) (expenses.Expenses, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var expenseList expenses.Expenses
	if err := json.Unmarshal(data, &expenseList); err != nil {
		return nil, err
	}
	return expenseList, nil
}

func _getExpenseFilePath() (string, error) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Format current month
	currentMonth := time.Now().Format("January_2006") // e.g., "October_2023"

	// Construct the directory path
	dirPath := fmt.Sprintf("%s/budget_logging", homeDir)

	// Check if directory exists; if not, create it
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			return "", err
		}
	}

	// Construct the full file path
	filePath := fmt.Sprintf("%s/%s_spendings.json", dirPath, currentMonth)

	return filePath, nil
}
