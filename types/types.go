package types

import (
	"time"
)

type Expense struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
}

type Expenses []Expense

func (e Expense) FormattedDate() string {
	return e.Date.Format("2006-01-02")
}

type Forecast struct {
	Amount   float64
	Category string
}

type MonthlyForecast []Forecast

type MonthlyBudgetInsights struct {
	Expenses []Expense  `json:"expenses"`
	Forecast []Forecast `json:"forecast"`
}
