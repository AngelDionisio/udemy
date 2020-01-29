package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// add number of items the wg should expect
	wg.Add(1)

	go func() {
		count("sheep")
		// signal routine has finished
		wg.Done()
	}()

	// block
	wg.Wait()
}

func count(thing string) {
	for i := 0; i <= 8; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
