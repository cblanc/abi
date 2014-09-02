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
	fmt.Println("Benchmarking: ", *url)
	fmt.Println("Number of requests: ", *n)

	var times []float64

	for i := 1; i < *n; i++ {
		t := start()
		request(*url)
		times = append(times, end(t))
	}

	total := float64(0)

	for _, time := range times {
		total += time
	}

	average := total / float64(*n) / 1000000

	fmt.Println("Benchmark Complete")
	fmt.Printf("Total Time: %.2f ms\n\n", total)
	fmt.Printf("Average Time: %.2f ms\n\n", average)

	os.Exit(0)
}
