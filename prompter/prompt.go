package prompter

import (
	"bufio"
	"fmt"
	"github.com/Seymour-creates/budget-cli/types"
	"os"
	"strconv"
	"strings"
	"time"
)

// promptUser takes a prompt and returns user input
func promptUser(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// PromptForExpenses prompts user for new expense data and returns new expense
func PromptForExpenses() types.Expenses {
	var collectedExpenses types.Expenses

	for {
		fmt.Println("Enter expense details or type 'exit' to finish:")

		amountStr := promptUser("Enter the amount spent (or 'exit' to stop): ")
		if amountStr == "exit" {
			break
		}
		// Convert amount string to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Invalid Amount. Please Enter a valid number.")
			continue
		}
		dateStr := promptUser("Enter the date (YYYY-MM-DD): ")
		// convert date string to time.Time
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Invalid Date. Please Enter a valid format.")
			continue
		}
		description := promptUser("Enter the description of the expense: ")
		category := promptUser("Enter the category of the expense (Bill, Debt, Entertainment, Groceries, Misc, Savings, Takeout): ")

		newExpense := types.Expense{
			Date:        date,
			Description: description,
			Amount:      amount,
			Category:    category,
		}

		collectedExpenses = append(collectedExpenses, newExpense)
	}

	return collectedExpenses
}

// PromptForecastReport prompts user for expenditure estimates for the current month
func PromptForecastReport() types.MonthlyForecast {
	var financialForecast types.MonthlyForecast

	for {
		fmt.Println("Enter forecasted expenses for the upcoming month... \ntype 'exit' to finish:")

		amountStr := promptUser("Enter the forecasted amount (or 'exit' to stop): ")
		if amountStr == "exit" {
			break
		}

		// Convert amount string to float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Invalid Amount. Please Enter a valid number.")
			continue
		}

		category := promptUser("Enter the category of the forecasted expense (Bill, Debt, Entertainment, Groceries, Misc, Savings, Takeout): ")

		forecast := types.Forecast{
			Amount:   amount,
			Category: category,
		}

		financialForecast = append(financialForecast, forecast)
	}

	return financialForecast
}
