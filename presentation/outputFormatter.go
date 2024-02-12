package presentation

import (
	"fmt"
	"github.com/Seymour-creates/budget-cli/finance"
	"strings"
	"time"
)

type FinanceDisplay struct {
}

func NewFinanceDisplay() *FinanceDisplay {
	return &FinanceDisplay{}
}

// Expenses formats and displays expenses within a specified date range and calculates the total.
func (fd *FinanceDisplay) Expenses(expenses finance.Expenses, fromDate, toDate time.Time) float64 {
	var total float64
	fmt.Printf("\n%-15s %-50s %-10s %-15s\n",
		centerString("Date", 15),
		centerString("Description", 50),
		"Amount",
		centerString("Category", 15),
	)
	fmt.Println(strings.Repeat("-", 90))
	for _, expense := range expenses {
		if (expense.Date.After(fromDate) || expense.Date.Equal(fromDate)) && (expense.Date.Before(toDate) || expense.Date.Equal(toDate)) {
			fmt.Printf("%-15s %-50s %-10.2f %-15s\n",
				centerString(expense.Date.Format("2006-01-02"), 15), // Directly using time.Time Format method
				centerString(expense.Description, 50),
				expense.Amount,
				centerString(expense.Category, 15),
			)
			total += expense.Amount
		}
	}
	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf("\n%-65s %-10.2f\n", "Total", total)
	return total
}

// BarChart generates and displays a bar chart comparing forecasted and actual expenses.
func (fd *FinanceDisplay) BarChart(forecast, expenses map[string]float64) {
	const maxBarLength = 50
	for category, forecastAmount := range forecast {
		expenseAmount := expenses[category] // Assumes 0 if not found
		percentageSpent := 100 * expenseAmount / forecastAmount
		barLength := int(percentageSpent / 100 * maxBarLength)
		bar := strings.Repeat("=", barLength) + strings.Repeat(" ", maxBarLength-barLength)
		fmt.Printf("%-10s [%s] %.2f%%\n", category, bar, percentageSpent)
	}
}

func centerString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	leftPadding := (width - len(s)) / 2
	rightPadding := width - len(s) - leftPadding
	return fmt.Sprintf("%*s%s%*s", leftPadding, "", s, rightPadding, "")
}
