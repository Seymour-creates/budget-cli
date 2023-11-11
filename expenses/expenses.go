package expenses

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	Date        time.Time
	Description string
	Amount      float64
	Category    string
}

type Forecast struct {
	Amount   float64
	Category string
}

type MonthlyForecast []Forecast

type Expenses []Expense

// FormattedDate returns a YYYY-MM-DD format of expense date
func (e Expense) FormattedDate() string {
	return e.Date.Format("2006-01-02")
}

func (e *Expenses) DisplayExpensesAndTotal(fromDate, toDate time.Time) float64 {
	var total float64
	var expensesToShow Expenses

	if fromDate.IsZero() && toDate.IsZero() {
		now := time.Now()
		fromDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		toDate = time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC)
	}
	fmt.Printf("\n%-15s %-50s %-10s %-15s\n",
		_centerString("Date", 15),
		_centerString("Description", 50),
		"Amount",
		_centerString("Category", 15))
	totalWidth := 15 + 50 + 10 + 15 + (3 * 3)
	fmt.Println(strings.Repeat("-", totalWidth))
	for _, expense := range *e {
		if (expense.Date.After(fromDate) || expense.Date.Equal(fromDate)) &&
			(expense.Date.Before(toDate) || expense.Date.Equal(toDate)) {
			total += expense.Amount
			expensesToShow = append(expensesToShow, expense)
			fmt.Printf("%-15s %-50s %-10.2f %-15s\n",
				_centerString(expense.FormattedDate(), 15),
				_centerString(expense.Description, 50),
				expense.Amount,
				_centerString(expense.Category, 15))
		}
	}

	fmt.Println(strings.Repeat("-", totalWidth))
	fmt.Printf("\n%-45s %-10.2f\n", "Total", total)
	fmt.Println()

	return total
}

func _centerString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	leftPadding := (width - len(s)) / 2
	rightPadding := width - len(s) - leftPadding
	return fmt.Sprintf("%*s%s%*s", leftPadding, "", s, rightPadding, "")
}

func CompareForecastToExpenses(forecast MonthlyForecast, expenses Expenses) (map[string]float64, map[string]float64) {
	var billsForecast, entForecast, debtForecast, groceryForecast, savingsForecast, miscForecast, takeoutForecast float64
	var billsExpenses, entExpenses, debtExpenses, groceryExpenses, savingsExpenses, miscExpenses, takeoutExpenses float64
	for _, catCast := range forecast {
		switch catCast.Category {
		case "bill":
			billsForecast = catCast.Amount
		case "ent":
			entForecast = catCast.Amount
		case "debt":
			debtForecast = catCast.Amount
		case "misc":
			miscForecast = catCast.Amount
		case "grocery":
			groceryForecast = catCast.Amount
		case "saving":
			savingsForecast = catCast.Amount
		case "takeout":
			takeoutForecast = catCast.Amount
		}
	}

	for _, expense := range expenses {
		switch {
		case strings.HasPrefix(expense.Category, "bill"):
			billsExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "ent"):
			entExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "debt"):
			debtExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "misc"):
			miscExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "grocery"):
			groceryExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "saving"):
			savingsExpenses += expense.Amount
		case strings.HasPrefix(expense.Category, "takeout"):
			takeoutExpenses += expense.Amount
		default:
			miscExpenses += expense.Amount
		}
	}

	// Create a map to return the forecast amounts
	forecastTotals := map[string]float64{
		"bill":    billsForecast,
		"ent":     entForecast,
		"debt":    debtForecast,
		"grocery": groceryForecast,
		"saving":  savingsForecast,
		"misc":    miscForecast,
		"takeout": takeoutForecast,
	}

	// Create a map to return the expense amounts
	expenseTotals := map[string]float64{
		"bill":    billsExpenses,
		"ent":     entExpenses,
		"debt":    debtExpenses,
		"grocery": groceryExpenses,
		"saving":  savingsExpenses,
		"misc":    miscExpenses,
		"takeout": takeoutExpenses,
	}

	// Now both maps are used to return the data
	return forecastTotals, expenseTotals

}

func PrintBarChart(forecast map[string]float64, expenses map[string]float64) {
	const maxBarLength = 50
	for category, forecastAmount := range forecast {
		expenseAmount, exists := expenses[category]
		if !exists {
			expenseAmount = 0
		}
		percentageSpent := (expenseAmount / forecastAmount) * 100
		barLength := int(percentageSpent / 100 * maxBarLength)
		bar := strings.Repeat("=", barLength) + strings.Repeat(" ", maxBarLength-barLength)
		fmt.Printf("%-10s [%s] %.2f%%\n", category, bar, percentageSpent)
	}
}
