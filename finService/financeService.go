package finService

import (
	"fmt"
	"github.com/Seymour-creates/budget-cli/finance"
)

type FinanceService struct {
}

func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

func (fs *FinanceService) ExtractForecastAndExpense(forecast finance.MonthlyForecast, expenses finance.Expenses) (forecastTotals, expenseTotals map[string]float64) {
	forecastTotals = make(map[string]float64)
	expenseTotals = make(map[string]float64)

	if len(forecast) == 0 {
		fmt.Println("No forecast data available. Please create a monthly forecast using the add-forecast command.")
		return nil, nil
	}

	for _, f := range forecast {
		forecastTotals[f.Category] = f.Amount
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses data available.")
	} else {
		for _, e := range expenses {
			expenseTotals[e.Category] += e.Amount
		}
	}

	return forecastTotals, expenseTotals
}
