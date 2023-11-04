package main

import (
	"bufio"
	"fmt"
	"github.com/Seymour-creates/budget-cli/cli"
	"github.com/Seymour-creates/budget-cli/expenses"
	"github.com/Seymour-creates/budget-cli/storage"
	"os"
	"strings"
	"time"
)

// TODO: Expand to cobra for cli args..
//
//	export saving data to db
//	expose cli via rest api
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please make a decision of what you would like to do:" +
		"\n\t1: Add expense" +
		"\n\t2: Forecast Monthly Expenses" +
		"\n\t3: Generate Spending Summary" +
		"\n\t4: Compare current expenses to monthly forecast\n\t➡️")

	choice, _ := reader.ReadString('\n')
	switch strings.TrimSpace(choice) {
	case "1":
		exp := cli.PromptForExpenses()
		fmt.Println("Expenses to be added:", exp)
		monthlyExpenses := expenses.Expenses{}
		monthlyExpenses.AddExpense(exp...)
		if err := storage.UpdateExpenses(monthlyExpenses...); err != nil {
			fmt.Println("Error updating expenses:", err)
			return
		}
		fmt.Println("Updated monthly expenses:", monthlyExpenses)
	case "2":
		forecast := cli.PromptForecastReport()
		fmt.Println("forecast: ", forecast)
		if err := storage.UpdateForecast(forecast...); err != nil {
			fmt.Println("Error updating forecast:", err)
			return
		}
		fmt.Println("Completed Forecast:", forecast)
	case "3":
		loadedExpenses, err := storage.LoadExpensesFromJSON()
		if err != nil {
			fmt.Println("Error loading expenses:", err)
			return
		}
		total := loadedExpenses.DisplayExpensesAndTotal(time.Time{}, time.Time{})
		fmt.Println("Total of all expenses:", total)
	case "4":
		moneySpent, err := storage.LoadExpensesFromJSON()
		if err != nil {
			fmt.Println("Error loading expenses for comparison.")
			return
		}
		cashFlow, err := storage.LoadForecastFromJSON()
		if err != nil {
			fmt.Println("Error loading forecast data for comparison.")
			return
		}
		forecasted, expenditure := expenses.CompareForecastToExpenses(cashFlow, moneySpent)
		expenses.PrintBarChart(forecasted, expenditure)
	default:
		return
	}
}

/*
/budget-cli
|--/expenses
|  |-- expenses.go
|  |-- expenses_test.go
|--/forecast
|--/storage
|--/report
|--/utils
|-- main.go
|-- go.mod
|-- go. Sum
*/
