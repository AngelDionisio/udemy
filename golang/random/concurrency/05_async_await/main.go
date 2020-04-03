package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRunningTask(waitInSeconds int) <-chan int {
	r := make(chan int)

	go func() {
		defer close(r)

		// Simulate workload
		fmt.Printf("starting work that will take %d seconds...\n", waitInSeconds)

		// set seed to assure random numbers
		rand.Seed(time.Now().UnixNano())
		min, max := 1, 100

		time.Sleep(time.Second * time.Duration(waitInSeconds))
		r <- rand.Intn((max - min + 1) + min)
	}()

	return r
}

func main() {
	// await a single goroutine
	r := <-longRunningTask(3)
	fmt.Println(r)

	// await multiple goroutines
	aChan, bChan, cChan := longRunningTask(3), longRunningTask(4), longRunningTask(2)
	a, b, c := <-aChan, <-bChan, <-cChan

	fmt.Println(a, b, c)
}
