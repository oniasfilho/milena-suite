package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL endpoint to stress test")
	requests := flag.Int("requests", 10, "Number of requests to make")
	concurrency := flag.Int("concurrency", 1, "Number of concurrent goroutines")
	
	flag.Parse()
	
	if *url == "" {
		fmt.Println("Error: URL flag is required")
		fmt.Println("Usage: go run main.go --url <endpoint> [--requests <num>] [--concurrency <num>]")
		os.Exit(1)
	}
	
	fmt.Printf("Starting stress test on %s with %d requests using %d concurrent workers\n", *url, *requests, *concurrency)
	
	semaphore := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup
	
	results := make(chan string, *requests)
	
	startTime := time.Now()
	
	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go func(requestNum int) {
			defer wg.Done()
			
			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			
			// Make HTTP request
			resp, err := http.Get(*url)
			if err != nil {
				results <- "ERROR"
				return
			}
			defer resp.Body.Close()
			
			results <- resp.Status
		}(i)
	}
	
	wg.Wait()
	close(results)
	
	duration := time.Since(startTime)
	
	statusCounts := make(map[string]int)
	for result := range results {
		statusCounts[result]++
	}
	
	fmt.Printf("\n--- Results ---\n")
	fmt.Printf("Total requests: %d\n", *requests)
	fmt.Printf("Concurrency: %d\n", *concurrency)
	fmt.Printf("Total time: %v\n", duration)
	fmt.Printf("Requests per second: %.2f\n\n", float64(*requests)/duration.Seconds())
	
	fmt.Println("Status Code Summary:")
	for status, count := range statusCounts {
		fmt.Printf("  %s: %d requests\n", status, count)
	}
} 