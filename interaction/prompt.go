package interaction

import (
	"bufio"
	"fmt"
	"github.com/Seymour-creates/budget-cli/finance"
	"os"
	"strconv"
	"strings"
	"time"
)

type Prompt struct {
	reader *bufio.Reader
}

func NewPrompter() *Prompt {
	return &Prompt{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (p *Prompt) promptUser(message string) (string, error) {
	fmt.Println(message)
	input, err := p.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (p *Prompt) amount() (float64, error) {
	amountStr, err := p.promptUser("Enter the amount spent (or 'exit' to stop): ")
	if err != nil {
		return 0, err
	}
	if amountStr == "exit" {
		return 0, fmt.Errorf("exit")
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount. Please Enter a valid number.")
		return 0, err
	}
	return amount, nil
}

func (p *Prompt) date() (time.Time, error) {
	dateStr, err := p.promptUser("Enter the date (YYYY-MM-DD): ")
	if err != nil {
		return time.Time{}, err
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Invalid Date. Please Enter a valid format.")
		return time.Time{}, err
	}
	return date, nil
}

func (p *Prompt) CollectExpenses() finance.Expenses {
	var collectedExpenses finance.Expenses

	for {
		fmt.Println("Enter expense details or type 'exit' to finish:")

		amount, err := p.amount()
		if err != nil {
			if err.Error() == "exit" {
				break
			}
			fmt.Println("Invalid amount. Please Enter a valid number.")
			continue
		}

		date, err := p.date()
		if err != nil {
			fmt.Println("Invalid Date. Please Enter a valid format.")
			continue
		}

		description, _ := p.promptUser("Enter the description of the expense: ")
		category, _ := p.promptUser("Enter the category of the expense (Bill, Debt, Entertainment, Groceries, Misc, Savings, Takeout): ")

		newExpense := finance.Expense{
			Date:        date,
			Description: description,
			Amount:      amount,
			Category:    category,
		}

		collectedExpenses = append(collectedExpenses, newExpense)
	}

	return collectedExpenses
}

func (p *Prompt) CollectForecast() finance.MonthlyForecast {
	var financialForecast finance.MonthlyForecast

	for {
		fmt.Println("Enter forecasted expenses for the upcoming month... \ntype 'exit' to finish:")

		amount, err := p.amount()
		if err != nil {
			if err.Error() == "exit" {
				break
			}
			fmt.Println("Invalid amount. Please Enter a valid number.")
			continue
		}

		category, _ := p.promptUser("Enter the category of the forecasted expense (Bill, Debt, Entertainment, Groceries, Misc, Savings, Takeout): ")

		forecast := finance.Forecast{
			Amount:   amount,
			Category: category,
		}

		financialForecast = append(financialForecast, forecast)
	}

	return financialForecast
}
