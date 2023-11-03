package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/expenses"
	"os"
	"time"
)

// TODO: Create a way to edit forecast and expenses
// 	forecast and expense should be able to be filtered by ..
// 	possibly may need to serialize entries

func UpdateExpenses(newExpenses ...expenses.Expense) error {
	existingExpenses, err := LoadExpensesFromJSON()
	if err != nil {
		fmt.Println("Error loading existing expenses:", err)
		return err
	}

	combinedExpense := append(existingExpenses, newExpenses...)
	return SaveExpensesToJSON(combinedExpense)
}

func UpdateForecast(newForecast ...expenses.Forecast) error {
	existingForecast, err := LoadForecastFromJSON()
	if err != nil {
		fmt.Println("Error loading existing forecast:", err)
		return err
	}

	combinedForecast := append(existingForecast, newForecast...)
	return SaveForecastToJSON(combinedForecast)
}

func LoadExpensesFromJSON() (expenses.Expenses, error) {
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

func LoadForecastFromJSON() (expenses.MonthlyForecast, error) {
	filePath, err := _getForecastFilePath()
	if err != nil {
		fmt.Println("Error getting forecast file path:", err)
		return nil, err
	}

	data, err := readFileContents(filePath)
	if err != nil {
		fmt.Println("Error reading the forecast file:", err)
		return nil, err
	}

	if len(data) == 0 {
		return expenses.MonthlyForecast{}, nil
	}

	forecast, err := unmarshalForecast(data)
	if err != nil {
		fmt.Println("Error unmarshalling the forecast content:", err)
		return nil, err
	}

	return forecast, nil
}

func SaveExpensesToJSON(expenses expenses.Expenses) error {
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

func SaveForecastToJSON(forecast expenses.MonthlyForecast) error {
	filePath, err := _getForecastFilePath()
	if err != nil {
		fmt.Println("Error getting forecast file path:", err)
		return err
	}

	data, err := marshalForecast(forecast)
	if err != nil {
		fmt.Println("Error marshalling forecast:", err)
		return err
	}

	if err := writeToFile(filePath, data); err != nil {
		fmt.Println("Error writing the forecast to the file:", err)
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

func marshalForecast(forecast expenses.MonthlyForecast) ([]byte, error) {
	return json.Marshal(forecast)
}

func unmarshalForecast(data []byte) (expenses.MonthlyForecast, error) {
	var currentForecast expenses.MonthlyForecast
	if err := json.Unmarshal(data, &currentForecast); err != nil {
		return nil, err
	}
	return currentForecast, nil
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

func _getForecastFilePath() (string, error) {
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
	filePath := fmt.Sprintf("%s/%s_forecast.json", dirPath, currentMonth)

	return filePath, nil
}
