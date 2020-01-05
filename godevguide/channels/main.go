package main

import (
	"fmt"
	"net/http"
	"time"
)

// concurrent programs can schedule multiple units of work at time (like go routines)
// having the ability to switch between them. And parallelism means having two or more
// processes running at the exact same time.
// So you can have one core running concurrent multiple goroutines, each one can go to sleep
// on a blocking action, like an http request, and the scheduler cycles between them as they wake up
// to execute the remainder of the code in the functions. These actions take fractions of fractions of a second
// These run in PARALLEL with another core also running multiple concurrent go routines.
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkInfinite(link, c)
	}

	// when data comes out of this channel, log it
	// sending and receiving messages is a blocking action
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	// in this instance, each iteration of the loop will block until the channel receives a message to print
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	// this will run forever, however as each request is blocking it is not executing a large # of requests a second
	// 	go checkLink2(<-c, c)
	// }

	// for ranging on a channel, wait for a message to come out a channel, and assign it to l
	for l := range c {
		// function literal, equivalent to an IIFE in JS, immediatetly-invoked function expression
		// we do this as the main func (goroutine) might have already changed the value of l by time checklink tries to reference it
		// so as a result, as go passes arguments by value, we pass the value of l to the function literal, so that a new copy
		// is saved in memory
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkInfinite(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yup its up"
}

// with this version we create a "is service up" mechanism
// we can put this function in an infinite for loop which will keep calling this function again and again
func checkLinkInfinite(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
