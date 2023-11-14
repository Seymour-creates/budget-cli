package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{Use: "budget-prompter"}

	// Add expense command
	var cmdAddExpense = &cobra.Command{
		Use:   "add-expense",
		Short: "Add a new expense",
		Run:   addExpenseCmd,
	}

	// Forecast monthly expenses command
	var cmdForecast = &cobra.Command{
		Use:   "forecast",
		Short: "Forecast monthly types",
		Run:   monthlyForecastCmd,
	}

	// Generate spending summary command
	var cmdSummary = &cobra.Command{
		Use:   "summary",
		Short: "Generate a spending summary",
		Run:   summaryCmd,
	}

	// Compare expenses to forecast command
	var cmdCompare = &cobra.Command{
		Use:   "compare",
		Short: "Compare current types to monthly forecast",
		Run:   compareForecastToExpenseCmd,
	}

	// Adding the commands to the root command
	rootCmd.AddCommand(cmdAddExpense, cmdForecast, cmdSummary, cmdCompare)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
/budget-prompter
|--/types
|--/prompter
|--/client
|--/report
|--/utils
|-- main.go
|-- commands.go
|-- go.mod
|-- go. Sum
*/
