package main

import "fmt"

func main() {
	jobs := make(chan int, 70)
	results := make(chan int, 70)

	go worker(jobs, results)

	for i := 0; i < 70; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 70; i++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	a := 0
	b := 1
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
		// fmt.Printf("temp: %v, a: %v, b: %v\n", temp, a, b)
	}
	return a
}

func slowFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return slowFib(n-1) + slowFib(n-2)
	}
}
