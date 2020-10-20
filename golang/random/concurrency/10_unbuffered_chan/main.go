// This small program demostrates how to use an unbuffered channel
// to simulate a game of tennis between two goroutines.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create unbuffered channel
	ballCourtChannel := make(chan int)

	// set number of goroutines to wait-for
	wg.Add(2)

	go player("Nadal", ballCourtChannel, &wg)
	go player("Djokovic", ballCourtChannel, &wg)

	// start game by sending ball into channel
	ballCourtChannel <- 1

	// wait for game to finish
	wg.Wait()

}

// player simulates a person playing a game of tennis
func player(name string, ballChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// infinite loop until the ball is not in channel
	// signifying that the channel was closed as the other player "missed"
	for {
		ball, ok := <-ballChan
		if !ok {
			// if channel is closed we won
			fmt.Printf("Player %s won!\n", name)
			return
		}

		// pick randon number to check if player missed
		// if player missed, close chan, exit function
		n := rand.Intn(100)
		if n%14 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(ballChan)
			return
		}

		// increase the value of ball, as a hit
		fmt.Printf("Player %s hit ball # %d\n", name, ball)
		ball++

		ballChan <- ball
	}
}
