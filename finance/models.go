package finance

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

type Forecast struct {
	Amount   float64
	Category string
}

type MonthlyForecast []Forecast

type MonthlyBudgetInsights struct {
	Expenses []Expense  `json:"expenses"`
	Forecast []Forecast `json:"forecast"`
}

type Categories struct{}

func (c Categories) IsValid(category string) bool {
	validCategories := []string{"bill", "debt", "ent", "grocery", "misc", "saving", "takeout"}
	for _, v := range validCategories {
		if v == category {
			return true
		}
	}
	return false
}
