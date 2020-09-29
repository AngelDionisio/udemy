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

	checkURLs(links)
}

func checkURLs(list []string) {
	var wg sync.WaitGroup
	c := make(chan string)

	// launch Goroutine that will block, waiting for WaitGroup queue to be depleted
	// then close channel used to receive messages
	go func() {
		wg.Wait()
		close(c)
	}()

	for _, url := range list {
		// tell WaitGroup to add one item to wait for
		// could be replaced with wg.Add(len(list))
		wg.Add(1)
		go isUp(url, c, &wg)
	}

	for msg := range c {
		fmt.Println(msg)
	}
}

// isUp is to be used as a Goroutine for checking if an URL is up
// it accepts the URL, a channel to send messages to, and a pointer to a Waitgroup
// to keep sync of how many of these the caller should wait for
func isUp(url string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%v is down!\n", url)
		return
	}
	c <- fmt.Sprintf("%v is up!", url)
}
