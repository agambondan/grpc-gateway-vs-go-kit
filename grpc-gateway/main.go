package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type requestPayload struct {
	A float64 `json:"a,omitempty"`
	B float64 `json:"b,omitempty"`
	N int32   `json:"n,omitempty"`
}

func main() {
	endpoints := []struct {
		name   string
		url    string
		method string
		body   requestPayload
	}{
		{"Add", "http://localhost:9901/api/v1/calculator/add", "POST", requestPayload{A: 1, B: 2}},
		{"Subtract", "http://localhost:9901/api/v1/calculator/subtract", "POST", requestPayload{A: 5, B: 3}},
		{"Multiply", "http://localhost:9901/api/v1/calculator/multiply", "POST", requestPayload{A: 2, B: 4}},
		{"Divide", "http://localhost:9901/api/v1/calculator/divide", "POST", requestPayload{A: 10, B: 2}},
		{"Fibonacci", "http://localhost:9901/api/v1/calculator/fibonacci", "POST", requestPayload{N: 100}},
	}

	var wg sync.WaitGroup
	g, ctx := errgroup.WithContext(context.Background())

	for _, endpoint := range endpoints {
		time.Sleep(time.Second)
		for i := 0; i < 100; i++ { // 1000 requests per endpoint
			wg.Add(1)
			e := endpoint // avoid closure issue
			g.Go(func() error {
				defer wg.Done()
				return sendRequest(ctx, e.url, e.method, e.body)
			})
		}
	}

	// Wait for all requests to finish
	wg.Wait()
	if err := g.Wait(); err != nil {
		fmt.Printf("Stress test failed: %v\n", err)
	} else {
		fmt.Println("Stress test completed successfully.")
	}
}

func sendRequest(ctx context.Context, url, method string, body requestPayload) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(req, resp)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
