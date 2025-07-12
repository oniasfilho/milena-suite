package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the URL flag
	url := flag.String("url", "", "URL endpoint to stress test")
	
	// Parse command line flags
	flag.Parse()
	
	// Check if URL flag was provided
	if *url == "" {
		fmt.Println("Error: URL flag is required")
		fmt.Println("Usage: go run main.go -url <endpoint>")
		os.Exit(1)
	}
	
	// Print the stress test message
	fmt.Printf("I'll stress test the %s endpoint\n", *url)
} 