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
	rootCmd := &cobra.Command{Use: "budgetcli"}
	command := generateCommandController()
	// Adding the commands to the root command
	rootCmd.AddCommand(
		command.AddExpenseCmd(),
		command.AddForecastCmd(),
		command.ExpenseDetailsCmd(),
		command.ExpenseReportCmd(),
	)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateCommandController() *commands.Command {
	prompter := interaction.NewPrompter()
	service := finService.NewFinanceService()
	presenter := presentation.NewFinanceDisplay()
	client := xatClient.NewClient()
	command := commands.NewCommander(prompter, presenter, service, client)
	return command
}
