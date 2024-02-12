package finService

import "github.com/Seymour-creates/budget-cli/finance"

type FinanceService struct {
}

func NewFinanceService() *FinanceService {
	return &FinanceService{}
}

func (fs *FinanceService) ExtractForecastAndExpense(forecast finance.MonthlyForecast, expenses finance.Expenses) (forecastTotals, expenseTotals map[string]float64) {

	for _, f := range forecast {
		forecastTotals[f.Category] = f.Amount
	}

	for _, e := range expenses {
		expenseTotals[e.Category] += e.Amount
	}

	return forecastTotals, expenseTotals
}
