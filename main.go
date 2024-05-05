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
	url := flag.String("url", "", "URL of the service to be tested")
	requests := flag.Int("requests", 0, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of concurrent calls")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Error: invalid arguments. Make sure to provide --url, --requests, and --concurrency correctly.")
		os.Exit(1)
	}

	fmt.Printf("Starting load tests on %s with %d requests and %d concurrent calls\n", *url, *requests, *concurrency)

	startTime := time.Now()
	successfulRequests, statusCodeCount := doLoadTest(*url, *requests, *concurrency)
	totalTime := time.Since(startTime)

	fmt.Printf("\nLoad Test Report:\n")
	fmt.Printf("Total time spent: %s\n", totalTime)
	fmt.Printf("Total number of requests: %d\n", *requests)
	fmt.Printf("Number of requests with status 200 OK: %d\n", successfulRequests)
	fmt.Println("Status Code Distribution:")
	for code, count := range statusCodeCount {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}

func doLoadTest(url string, requests, concurrency int) (int, map[int]int) {
	var (
		wg               sync.WaitGroup
		successfulRequests int
		statusCodeCount     = make(map[int]int)
		requestsChan       = make(chan struct{}, concurrency)
	)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		requestsChan <- struct{}{}
		go func() {
			defer func() {
				<-requestsChan
				wg.Done()
			}()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error making request: %v\n", err)
				statusCodeCount[0]++
				return
			}
			defer resp.Body.Close()

			statusCodeCount[resp.StatusCode]++
			if resp.StatusCode == http.StatusOK {
				successfulRequests++
			}
		}()
	}

	wg.Wait()

	return successfulRequests, statusCodeCount
}
