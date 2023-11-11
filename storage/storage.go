package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Seymour-creates/budget-cli/expenses"
	"io"
	"log"
	"net/http"
	"time"
)

// TODO: Create a way to edit forecast and expenses
// 	forecast and expense should be able to be filtered by ..
// 	possibly may need to serialize entries

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

func PostExpense(expenses expenses.Expenses) error {
	expensesToPost, err := json.Marshal(expenses)
	if err != nil {
		log.Fatalf("error stringifying slice: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	success, err := makeHTTPReq(ctx, http.MethodPost, "http://localhost:3000/add_expense", headers, expensesToPost)
	if err != nil {
		log.Printf("Error making http request: %v", err)
	}
	fmt.Println("Response: ", string(success))
	return nil
}
