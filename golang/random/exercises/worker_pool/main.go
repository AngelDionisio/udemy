package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func worker(id int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("WorkerID: %d picking up job", id)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}

func main() {
	fmt.Println("Worker pools in Go")

	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
	}

	now := time.Now()
	numOfWorkers := 3
	processResults(urls, numOfWorkers)
	timeElapsed := time.Since(now)

	now2 := time.Now()
	numOfWorkers2 := 6
	processResults(urls, numOfWorkers2)
	timeElapsed2 := time.Since(now2)
	fmt.Printf("It took %v for job to complete with %d number of workers\n", timeElapsed, numOfWorkers)
	fmt.Printf("It took %v for job to complete with %d number of workers\n", timeElapsed2, numOfWorkers2)
}

func processResults(urls []string, numOfWorkers int) {
	// buffered job stream that can hold the len(urls) # of jobs
	jobs := make(chan Site, len(urls))
	// results stream can take len(urls) number of results
	results := make(chan Result, len(urls))

	// instantiating N number of workers to split work
	for workerID := 1; workerID <= numOfWorkers; workerID++ {
		go worker(workerID, jobs, results)
	}

	// send jobss into job stream
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs) // always close channels from the sending end, trying to send to a close channel, panics

	var counter int
	for i := 0; i < len(urls); i++ {
		result := <-results
		counter++
		log.Println(result)
	}

	fmt.Printf("lenOfJobs: %d, lenOfResults: %d\n", len(urls), counter)

	close(results)
}
