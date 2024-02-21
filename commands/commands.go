package commands

import (
	"context"
	"fmt"
	"github.com/Seymour-creates/budget-cli/finService"
	"github.com/Seymour-creates/budget-cli/interaction"
	"github.com/Seymour-creates/budget-cli/presentation"
	"github.com/Seymour-creates/budget-cli/xatClient"
	"github.com/spf13/cobra"
	"time"
)

type Command struct {
	prompter *interaction.Prompt
	present  *presentation.FinanceDisplay
	service  *finService.FinanceService
	client   *xatClient.Client
}

func NewCommander(p *interaction.Prompt, present *presentation.FinanceDisplay, service *finService.FinanceService, client *xatClient.Client) *Command {
	return &Command{
		prompter: p,
		present:  present,
		service:  service,
		client:   client,
	}
}

func (c *Command) AddExpenseCmd() *cobra.Command {
	ctx := context.Background()
	return &cobra.Command{
		Use:   "add-expense",
		Short: "Add expense",
		Run: func(cmd *cobra.Command, args []string) {
			exp := c.prompter.CollectExpenses()
			fmt.Println("Expenses to be added:", exp)
			if err := c.client.PostExpense(ctx, exp); err != nil {
				fmt.Println("Error posting types: ", err)
			}
		},
	}
}

func (c *Command) AddForecastCmd() *cobra.Command {
	ctx := context.Background()
	return &cobra.Command{
		Use:   "add-forecast",
		Short: "Monthly forecast",
		Run: func(cmd *cobra.Command, args []string) {
			forecast := c.prompter.CollectForecast()
			fmt.Println("Forecast: ", forecast)
			if err := c.client.PostForecast(ctx, forecast); err != nil {
				fmt.Println("Error post forecast: ", err)
			}
		},
	}
}

func (c *Command) ExpenseDetailsCmd() *cobra.Command {
	ctx := context.Background()
	return &cobra.Command{
		Use:   "detail-expenses",
		Short: "Generate a detailed expenses summary",
		Run: func(cmd *cobra.Command, args []string) {
			loadedExpenses, err := c.client.GetMonthExpenses(ctx)
			if err != nil {
				fmt.Println("Error getting expense data: ", err)
			}

			t := time.Now()
			startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
			endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

			total := c.present.Expenses(loadedExpenses, startOfMonth, endOfMonth)
			fmt.Println("Total of all expenses:", total)
		},
	}
}

func (c *Command) ExpenseReportCmd() *cobra.Command {
	ctx := context.Background()
	return &cobra.Command{
		Use:   "expense-report",
		Short: "Compare current expenses and purchases to monthly forecast",
		Run: func(cmd *cobra.Command, args []string) {
			expenses, forecast, err := c.client.GetForecastAndExpense(ctx)
			if err != nil {
				fmt.Println("Error getting expenses or expense data: ", err)
			}
			forecasted, expenditure := c.service.ExtractForecastAndExpense(forecast, expenses)
			c.present.BarChart(forecasted, expenditure)
		},
	}
}

func (c *Command) RefreshExpenseDataFromPlaidCmd() *cobra.Command {
	ctx := context.Background()
	return &cobra.Command{
		Use:   "update-expenses",
		Short: "Fetches latest expenses for the month from plaid, and updates db with expenses",
		Run: func(cmd *cobra.Command, args []string) {
			err := c.client.RefreshExpenseDataViaPlaid(ctx)
			if err != nil {
				fmt.Println("error fetching expense data from plaid: ", err)
			}
		},
	}
}
