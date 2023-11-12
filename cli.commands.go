package main

import (
	"fmt"
	"github.com/Seymour-creates/budget-cli/client"
	"github.com/Seymour-creates/budget-cli/prompter"
	"github.com/Seymour-creates/budget-cli/types"
	"github.com/spf13/cobra"
	"time"
)

func addExpenseCmd(cmd *cobra.Command, args []string) {
	exp := prompter.PromptForExpenses()
	fmt.Println("Expenses to be added:", exp)
	if err := client.PostExpense(exp); err != nil {
		fmt.Println("Error posting types: ", err)
	}
}

func monthlyForecastCmd(cmd *cobra.Command, args []string) {
	forecast := prompter.PromptForecastReport()
	fmt.Println("Forecast: ", forecast)
	if err := client.PostForecast(forecast); err != nil {
		fmt.Println("Error post forecast: ", err)
	}
}

func summaryCmd(cmd *cobra.Command, args []string) {
	loadedExpenses, err := client.GetMonthlySummary()
	if err != nil {
		fmt.Println("Error getting expense data: ", err)
	}

	total := loadedExpenses.DisplayExpensesAndTotal(time.Time{}, time.Time{})
	fmt.Println("Total of all expenses:", total)
}

func compareForecastToExpenseCmd(cmd *cobra.Command, args []string) {
	forecast, cashOut, err := client.GetCompareForecastToExpense()
	if err != nil {
		fmt.Println("Error getting forecast or expense data: ", err)
	}
	forecasted, expenditure := types.CompareForecastToExpenses(cashOut, forecast)
	types.PrintBarChart(forecasted, expenditure)
}
