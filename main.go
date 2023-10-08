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
		fmt.Println(exp)
		monthlyExpenses := expenses.Expenses{}
		monthlyExpenses.AddExpense(exp...)
		if err := storage.UpdateExpenses(monthlyExpenses...); err != nil {
			return
		}
		fmt.Println("monthly expenses", monthlyExpenses)
	case "2": //
	case "3":
		loadedExpenses, err := storage.LoadExpenses()
		if err != nil {
			fmt.Println("Error loading expenses: ", err)
		}
		total := loadedExpenses.TotalExpense(time.Time{}, time.Time{})
		fmt.Println("Total of all expenses: ", total)
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
