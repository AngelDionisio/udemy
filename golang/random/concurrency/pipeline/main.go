package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := square(c)

	// consume all values in channel
	for res := range out {
		fmt.Println(res)
	}

	// these channels can be composed together
	for n := range square(square(gen(2, 3))) {
		fmt.Println(n)
	}
}

// gen converts a list of integers into a channel that emits the integers on the list
// starts a goroutine that sends the integers on the channel and closes the channel when all
// the values have been sent.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

// the second stage, square, receives the integers from a channel and returns a channel that emits
// the square of each integer received.
// after the inbound channel has been closed, and this stage has sent all values downstream, it closes
// the outbound channel
func square(inputChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range inputChan {
			out <- n * n
		}
		close(out)
	}()
	return out
}
