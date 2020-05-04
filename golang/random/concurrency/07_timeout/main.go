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

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
