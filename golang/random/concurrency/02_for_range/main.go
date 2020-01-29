package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	// for - range on a channel is syntactic sugar for receiving messages from a given channel
	// for as long as it is open. One could do this manually with example commented below
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
