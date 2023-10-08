package expenses

import (
	"time"
)

type Expense struct {
	Date        time.Time
	Description string
	Amount      float64
	Category    string
}

type Expenses []Expense

// AddExpense adds new expense to current month expenses
func (e *Expenses) AddExpense(expensesToAdd ...Expense) {
	for _, expense := range expensesToAdd {
		*e = append(*e, expense)
	}
}

// FormattedDate returns a YYYY-MM-DD format of expense date
func (e Expense) FormattedDate() string {
	return e.Date.Format("2006-01-02")
}

// TotalExpense calculates the total expense for a given period
func (e *Expenses) TotalExpense(fromDate, toDate time.Time) float64 {
	var total float64

	// If no dates are provided, sum all expenses
	if fromDate.IsZero() && toDate.IsZero() {
		for _, expense := range *e {
			total += expense.Amount
		}
		return total
	}

	// If only fromDate is provided, sum expenses from that date onwards
	if !fromDate.IsZero() && toDate.IsZero() {
		for _, expense := range *e {
			if expense.Date.After(fromDate) || expense.Date.Equal(fromDate) {
				total += expense.Amount
			}
		}
		return total
	}

	// If only toDate is provided, sum expenses until that date
	if fromDate.IsZero() && !toDate.IsZero() {
		for _, expense := range *e {
			if expense.Date.Before(toDate) || expense.Date.Equal(toDate) {
				total += expense.Amount
			}
		}
		return total
	}

	// If both dates are provided, sum expenses in that range
	for _, expense := range *e {
		if (expense.Date.After(fromDate) || expense.Date.Equal(fromDate)) &&
			(expense.Date.Before(toDate) || expense.Date.Equal(toDate)) {
			total += expense.Amount
		}
	}

	return total
}

//
//
//package expenses
//
//import (
//    "time"
//    // other necessary imports
//)
//
//// Expense represents a single expense entry
//type Expense struct {
//    Date        time.Time
//    Description string
//    Amount      float64
//    Category    string
//}
//
//// List of Expenses
//type Expenses []Expense
//
//// AddExpense appends a new expense to the list
//func (e *Expenses) AddExpense(expense Expense) {
//    *e = append(*e, expense)
//    // Optionally: Save to storage immediately after adding
//}
//
//// DeleteExpense removes an expense based on some criteria, e.g., a unique ID or date
//func (e *Expenses) DeleteExpense(/* criteria */) {
//// Logic to delete an expense
//}
//
//// UpdateExpense modifies an existing expense
//func (e *Expenses) UpdateExpense(expense Expense) {
//	// Logic to update an expense
//}
//
//// TotalExpense calculates the total expense for a given period
//func (e *Expenses) TotalExpense(fromDate, toDate time.Time) float64 {
//	// Logic to sum expenses in the given date range
//}
//
//// Other utility functions related to expenses
