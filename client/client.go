package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/types"
	"io"
	"log"
	"net/http"
	"time"
)

func makeHTTPReq(context context.Context, method, route string, headers map[string]string, body []byte) ([]byte, error) {

	req, err := http.NewRequestWithContext(context, method, route, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
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

func PostExpense(expenses types.Expenses) error {
	expensesToPost, err := json.Marshal(expenses)
	if err != nil {
		log.Fatalf("error stringifying slice: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	success, err := makeHTTPReq(ctx, http.MethodPost, "http://localhost:3000/post_expense", headers, expensesToPost)
	if err != nil {
		fmt.Errorf("error making http request: %v", err)
	}
	fmt.Println("Response: ", string(success))
	return nil
}

func PostForecast(forecast types.MonthlyForecast) error {
	forecastToPost, err := json.Marshal(forecast)
	if err != nil {
		log.Printf("Error stringifying slice: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	success, err := makeHTTPReq(ctx, http.MethodPost, "http://localhost:3000/post_forecast", headers, forecastToPost)
	if err != nil {
		fmt.Errorf("error making http request: %v", err)
	}
	fmt.Println("Response: ", success)
	return nil
}

func GetCompareForecastToExpense() (expenses types.Expenses, forecast types.MonthlyForecast, err error) {
	var _expense types.Expense
	var _forecast types.Forecast
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := makeHTTPReq(ctx, http.MethodGet, "http://localhost:3000/get_compare", headers, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error making http request: %v", err)
	}
	fmt.Println("Response data: ", resp)
	return types.Expenses{_expense}, types.MonthlyForecast{_forecast}, nil
}

func GetMonthlySummary() (types.Expenses, error) {
	var expenses types.Expenses

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	data, err := makeHTTPReq(ctx, http.MethodGet, "http://localhost:3000/get_summary", headers, nil)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, fmt.Errorf("error parsing response data %v.\nError: %v", data, err)
	}

	fmt.Println("Response data: ", data, "Expenses: ", expenses)

	return expenses, err
}
