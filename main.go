package main

import (
	"fmt"
	"github.com/Seymour-creates/budget-cli/commands"
	"github.com/Seymour-creates/budget-cli/finService"
	"github.com/Seymour-creates/budget-cli/interaction"
	"github.com/Seymour-creates/budget-cli/presentation"
	"github.com/Seymour-creates/budget-cli/xatClient"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rootCmd := &cobra.Command{Use: "budget-prompter"}
	prompter := interaction.NewPrompter()
	service := finService.NewFinanceService()
	presenter := presentation.NewFinanceDisplay()
	client := xatClient.NewClient()
	command := commands.NewCommand(prompter, presenter, service, client)
	// Adding the commands to the root command
	rootCmd.AddCommand(
		command.AddExpenseCmd(),
		command.MonthlyForecastCmd(),
		command.SummaryCmd(),
		command.CompareForecastToExpenseCmd(),
	)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
