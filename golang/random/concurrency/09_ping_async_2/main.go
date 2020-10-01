package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	pingURLs(links)
}

func pingURL(url string) string {
	_, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%v is down!", url)
	}
	return fmt.Sprintf("%v is up!", url)
}

func pingURLs(listOfUrls []string) {
	var wg sync.WaitGroup
	c := make(chan string)

	// launch Goroutine to wait for all tasks in WaitGroup to complete
	// then close channel
	go func() {
		wg.Wait()
		close(c)
	}()

	// create a Goroutine for each request, send results through channel
	for _, url := range listOfUrls {
		wg.Add(1)
		go func(lnk string) {
			defer wg.Done()
			c <- pingURL(lnk)
		}(url)
	}

	// extract messages from channel
	for msg := range c {
		fmt.Println(msg)
	}
}
