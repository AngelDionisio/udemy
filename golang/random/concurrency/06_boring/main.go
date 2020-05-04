package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring returns a receive only channel of strings
// sends message to channel on an infinite loop, so it will be ready to send
// as long as there is a channel waiting to receive.
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // Return channel to caller
}

// fanIn: multiple channels, funneling messages through another (one) channel
// lets any of the two channels communicate whenever they are ready
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

// fanInSelect rewrite of original fanIn, only one goRoutine is needed.
func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// When the main function executes <-c, it will wait for a value to be sent.
// Similarly, the boring function sends values to the channel only when a receiver is ready
// Both sender and receiver must be ready to play their part in the communication.
// Don't communicate by sharing memory, share memory by communicating
// Channels are first class values, just like strings and integers
func main() {
	joe := boring("Joe") // Function returning a channel
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	c := fanIn(boring("James"), boring("Miss Moneypenny"))
	for y := 0; y < 5; y++ {
		fmt.Println(<-c)
	}

	fmt.Println("You are both boring. I am leaving.")

	c2 := fanInSelect(boring("foo"), boring("bar"))
	for z := 0; z < 5; z++ {
		fmt.Println(<-c2)
	}
}
