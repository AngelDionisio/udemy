// This sample demostrates how to use a buffered channel
// to work on multiple tasks with a predefined number
// of goroutines

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberOfGoroutines = 4  // number of goroutines to use
	taskLoad           = 20 // number of tasks to process
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create a buffered channel that can hold 'taskLoad' number of items
	tasksChannel := make(chan string, taskLoad)

	wg.Add(numberOfGoroutines)

	for gr := 1; gr <= numberOfGoroutines; gr++ {
		go worker(tasksChannel, gr, &wg)
	}

	for job := 1; job <= taskLoad; job++ {
		tasksChannel <- fmt.Sprintf("Task: %v", job)
	}

	close(tasksChannel)

	wg.Wait()
}

func worker(tasks chan string, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// infinite loop, which allows worker to keep taking tasks from the
	// channel until there is nothing left to pick
	for {
		task, ok := <-tasks
		if !ok {
			// If there are no tasks in the channel, close worker
			fmt.Printf("No more tasks to process, worker: %d shutting down\n", workerID)
			return
		}

		// Log that worker is picking up a task
		fmt.Printf("Worker: %d is picking up task: %s\n", workerID, task)

		// sleep for random amount of time to simulate work
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// log that worker finished work
		fmt.Printf("Worker: %v finished task: %v in: %v\n", workerID, task, time.Duration(sleep)*time.Millisecond)
	}
}
