package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
	"strconv"
	"time"
)

type abiConfig struct {
	url      string
	requests int
}

type abiResult struct {
	totalTime           float64
	averageTime         float64
	successfulResponses int
}

func start() time.Time {
	return time.Now()
}

func end(t time.Time) float64 {
	return float64(time.Since(t).Nanoseconds())
}

func toSeconds(time float64) string {
	return strconv.FormatFloat(time/1000000000, 'f', 6, 64)
}

func toMs(time float64) string {
	return strconv.FormatFloat(time/1000000, 'f', 6, 64)
}

func toPercent(num int, den int) string {
	return strconv.FormatFloat(100*float64(num)/float64(den), 'f', 6, 64)
}

func request(url string) (*http.Response, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	return response, err
}

func extractConfig() abiConfig {
	var config abiConfig
	config.url = *flag.String("url", "http://localhost:9999/", "URL to request")
	config.requests = *flag.Int("n", 1000, "URL to request")
	flag.Parse()
	return config
}

func digestResults(config abiConfig, times []float64, successfulResponses int) abiResult {
	var result abiResult

	total := float64(0)
	for _, time := range times {
		total += time
	}
	result.totalTime = total

	result.averageTime = total / float64(config.requests)

	result.successfulResponses = successfulResponses

	return result
}

func presentResults(result abiResult) {

	data := [][]string{
		[]string{"Total Time (seconds)", toSeconds(result.totalTime)},
		[]string{"Average Time (ms)", toMs(result.averageTime)},
	}

	table := tablewriter.NewWriter(os.Stdout)

	for _, v := range data {
		table.Append(v)
	}

	fmt.Println("Benchmark Complete")

	table.Render()
}

func main() {
	config := extractConfig()

	fmt.Printf("Benchmarking: %s\n", config.url)
	fmt.Printf("Number of requests: %d\n\n", config.requests)

	var times []float64
	successfulResponses := 0

	for i := 0; i < config.requests; i++ {
		t := start()
		response, err := request(config.url)
		if err == nil && response.StatusCode == 200 {
			successfulResponses++
			times = append(times, end(t))
		}
	}

	result := digestResults(config, times, successfulResponses)
	presentResults(result)
}
