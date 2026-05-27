package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const requestTimeout = 10 * time.Second

type Result struct {
	URL        string
	StatusCode int
	ResponseMs int64
	Err        error
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: apicheck <file>")
		fmt.Println("Example: apicheck urls.txt")
		os.Exit(1)
	}

	filePath := os.Args[1]

	urls, err := readURLs(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(urls) == 0 {
		fmt.Println("No URLs found in file.")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	results := make(chan Result)

	client := &http.Client{
		Timeout: requestTimeout,
	}

	for _, endpoint := range urls {
		wg.Add(1)

		go func(endpoint string) {
			defer wg.Done()

			result := checkAPI(client, endpoint)
			results <- result
		}(endpoint)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		printResult(result)
	}
}

func readURLs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func checkAPI(client *http.Client, endpoint string) Result {
	parsedURL, err := url.ParseRequestURI(endpoint)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return Result{
			URL: endpoint,
			Err: fmt.Errorf("invalid URL"),
		}
	}

	start := time.Now()

	resp, err := client.Get(endpoint)
	if err != nil {
		return Result{
			URL: endpoint,
			Err: err,
		}
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Milliseconds()

	return Result{
		URL:        endpoint,
		StatusCode: resp.StatusCode,
		ResponseMs: elapsed,
	}
}

func printResult(result Result) {
	if result.Err != nil {
		fmt.Printf("✗ %s ERROR: %v\n", result.URL, result.Err)
		return
	}

	if result.StatusCode >= 200 && result.StatusCode < 400 {
		fmt.Printf("✓ %s %d (%dms)\n",
			result.URL,
			result.StatusCode,
			result.ResponseMs,
		)
	} else {
		fmt.Printf("⚠ %s %d (%dms)\n",
			result.URL,
			result.StatusCode,
			result.ResponseMs,
		)
	}
}