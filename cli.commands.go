package main

import (
	"fmt"
	"github.com/Seymour-creates/budget-cli/expenses"
	"github.com/Seymour-creates/budget-cli/prompter"
	"github.com/Seymour-creates/budget-cli/storage"
	"github.com/spf13/cobra"
	"time"
)

func addExpenseCmd(cmd *cobra.Command, args []string) {
	exp := prompter.PromptForExpenses()
	fmt.Println("Expenses to be added:", exp)
	if err := storage.PostExpense(exp); err != nil {
		fmt.Println("Error posting expenses: ", err)
	}
}

func monthlyForecastCmd(cmd *cobra.Command, args []string) {
	forecast := prompter.PromptForecastReport()
	fmt.Println("Forecast: ", forecast)
	// post forecast to db
}

func summaryCmd(cmd *cobra.Command, args []string) {
	loadedExpenses := expenses.Expenses{}
	total := loadedExpenses.DisplayExpensesAndTotal(time.Time{}, time.Time{})
	fmt.Println("Total of all expenses:", total)
}

func compareMonthToForecastCmd(cmd *cobra.Command, args []string) {
	cashFlow := expenses.MonthlyForecast{}
	moneySpent := expenses.Expenses{}
	forecasted, expenditure := expenses.CompareForecastToExpenses(cashFlow, moneySpent)
	expenses.PrintBarChart(forecasted, expenditure)
}
