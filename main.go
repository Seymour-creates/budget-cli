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

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please make a decision of what you would like to do:\n1: Add expense\n2: Forecast Monthly Expenses\n3: Generate Spending Summary\n")

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
