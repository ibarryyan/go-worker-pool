package main

import (
	"fmt"
	"github.com/ibarryyan/go-workers-pool/internal"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://bing.com",
		"https://apple.com",
	}

	task := func(url string) (string, error) {
		return fmt.Sprintf("task ... %s", url), nil
	}

	workerPool := internal.NewWorkerPool(3, task) // Create a worker pool with 3 workers
	workerPool.Start()

	// Submit the URLs to the worker pool for processing
	for _, url := range urls {
		workerPool.Submit(url)
	}

	// Collect the results and handle any errors
	for i := 0; i < len(urls); i++ {
		result := workerPool.GetResult()
		if result.err != nil {
			fmt.Printf("Worker ID: %d, URL: %s, Error: %v\n", result.workerID, result.url, result.err)
		} else {
			fmt.Printf("Worker ID: %d, URL: %s, Data: %s\n", result.workerID, result.url, result.data)
			// Save the extracted data to the database or process it further
			saveToDatabase(result.url, result.data)
		}
	}
}

// function to save the data to the database, replace this with actual database logic
func saveToDatabase(url, data string) {
	// Save the data to the database
}
