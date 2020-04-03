package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRunningTask(waitInSeconds int) <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// Simulate workload
		fmt.Println("starting work...")
		seed := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(seed)

		time.Sleep(time.Second * time.Duration(waitInSeconds))
		r <- r1.Int31n(100)
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
