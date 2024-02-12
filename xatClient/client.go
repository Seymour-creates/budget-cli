package xatClient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/finance"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	BaseURL = "https://xat.ngrok.app"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) makeHTTPReq(ctx context.Context, method, route string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, BaseURL+route, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("AUTH_USER"), os.Getenv("AUTH_USER_PASS"))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to url %v: %v", route, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response data for resp %v: %v", resp, err)
	}

	return data, nil
}

func (c *Client) PostExpense(ctx context.Context, expenses finance.Expenses) error {
	expensesToPost, err := json.Marshal(expenses)
	if err != nil {
		return fmt.Errorf("error stringifying slice: %v", err)
	}

	success, err := c.makeHTTPReq(ctx, http.MethodPost, "/post_expense", expensesToPost)
	if err != nil {
		return fmt.Errorf("error making http request: %v", err)
	}
	fmt.Println("Response: ", string(success))
	return nil
}

func (c *Client) PostForecast(ctx context.Context, forecast finance.MonthlyForecast) error {
	forecastToPost, err := json.Marshal(forecast)
	if err != nil {
		return fmt.Errorf("error stringifying slice: %v", err)
	}

	success, err := c.makeHTTPReq(ctx, http.MethodPost, "/post_forecast", forecastToPost)
	if err != nil {
		return fmt.Errorf("error making http request: %v", err)
	}
	fmt.Println("Response: ", string(success))
	return nil
}

func (c *Client) GetForecastAndExpense(ctx context.Context) (expenses finance.Expenses, forecast finance.MonthlyForecast, err error) {
	resp, err := c.makeHTTPReq(ctx, http.MethodGet, "/get_compare", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error making http request: %v", err)
	}
	var response finance.MonthlyBudgetInsights
	if err = json.Unmarshal(resp, &response); err != nil {
		return nil, nil, fmt.Errorf("error converting response: %v", err)
	}
	return response.Expenses, response.Forecast, nil
}

func (c *Client) GetMonthExpenses(ctx context.Context) (finance.Expenses, error) {
	var expenses finance.Expenses

	data, err := c.makeHTTPReq(ctx, http.MethodGet, "/get_expenses", nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, fmt.Errorf("error parsing response data %v.\nError: %v", string(data), err)
	}

	return expenses, err
}
