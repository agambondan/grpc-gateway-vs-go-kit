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
	// it shows your line code while print
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
		{"Fibonacci", "http://localhost:9901/api/v1/calculator/fibonacci/1000000", "GET", requestPayload{N: 1000}},
	}

	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	g, ctx := errgroup.WithContext(ctx)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, endpoint := range endpoints {
			wg.Add(1)
			go func(e struct {
				name   string
				url    string
				method string
				body   requestPayload
			},
			) {
				defer wg.Done()
				for i := 0; i < 10000; i++ {
					wg.Add(1)
					g.Go(func() error {
						defer wg.Done()
						return sendRequest(ctx, e.url, e.method, e.body)
					})
				}
			}(endpoint)
		}
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for _, endpoint := range endpoints {
	// 		for i := 0; i < 10000; i++ {
	// 			wg.Add(1)
	// 			e := endpoint
	// 			g.Go(func() error {
	// 				defer wg.Done()
	// 				return sendRequest(ctx, e.url, e.method, e.body)
	// 			})
	// 		}
	// 	}
	// }()

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
		log.Println(err)
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	buffer := bytes.NewBuffer(bodyBytes)
	if method == http.MethodGet {
		buffer = &bytes.Buffer{}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buffer)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(req, resp)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
