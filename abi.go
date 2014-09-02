package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func start() time.Time {
	return time.Now()
}

func end(t time.Time) float64 {
	return float64(time.Since(t).Nanoseconds())
}

func request(url string) (*http.Response, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	return response, err
}

func main() {
	url := flag.String("url", "http://localhost:9999/", "URL to request")
	n := flag.Int("n", 1000, "URL to request")
	flag.Parse()

	fmt.Printf("Benchmarking: %s\n", *url)
	fmt.Printf("Number of requests: %d\n\n", *n)

	var times []float64
	successfulResponses := 0

	for i := 0; i < *n; i++ {
		t := start()
		response, err := request(*url)
		if err == nil && response.StatusCode == 200 {
			successfulResponses++
			times = append(times, end(t))
		}
	}

	total := float64(0)

	for _, time := range times {
		total += time
	}

	average := total / float64(*n)

	fmt.Println("Benchmark Complete")
	fmt.Printf("Total Time: %.2f s\n", total/1000000000)
	fmt.Printf("Average Time: %.2f ms\n", average/1000000)
	fmt.Printf("Requests Succeeded: %.1f%%", 100*float64(successfulResponses)/float64(*n))
	fmt.Printf("\n\n")

	os.Exit(0)
}
