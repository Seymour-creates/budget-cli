package types

type Forecast struct {
	Amount   float64
	Category string
}

type MonthlyForecast []Forecast
