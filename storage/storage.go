package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/expenses"
	"os"
	"time"
)

func UpdateExpenses(newExpenses ...expenses.Expense) error {
	existingExpenses, err := LoadExpenses()
	if err != nil {
		fmt.Println("Error loading existing expenses:", err)
		return err
	}

	combinedExpense := append(existingExpenses, newExpenses...)
	return SaveExpenses(combinedExpense)
}

func LoadExpenses() (expenses.Expenses, error) {
	filePath, err := _getExpenseFilePath()
	if err != nil {
		fmt.Println("Error getting expense file path:", err)
		return nil, err
	}

	data, err := readFileContents(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}

	if bytes.Equal(data, []byte{}) {
		return expenses.Expenses{}, nil
	}

	expenseList, err := unmarshalExpenses(data)
	if err != nil {
		fmt.Println("Error unmarshalling the file content:", err)
		return nil, err
	}

	return expenseList, nil
}

func SaveExpenses(expenses expenses.Expenses) error {
	filePath, err := _getExpenseFilePath()
	if err != nil {
		fmt.Println("Error getting expense file path:", err)
		return err
	}

	data, err := marshalExpenses(expenses)
	if err != nil {
		fmt.Println("Error marshalling expenses:", err)
		return err
	}

	if err := writeToFile(filePath, data); err != nil {
		fmt.Println("Error writing to the file:", err)
		return err
	}

	return nil
}

func readFileContents(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return data, nil
}

func writeToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func marshalExpenses(expenses expenses.Expenses) ([]byte, error) {
	return json.Marshal(expenses)
}

func unmarshalExpenses(data []byte) (expenses.Expenses, error) {
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
