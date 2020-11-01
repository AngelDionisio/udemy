package main

import (
	"log"
	"sync"
	"time"

	"github.com/angeldionisio/udemy/golang/random/concurrency/12_worker_set_routines/work"
)

var names = []string{
	"angel",
	"julio",
	"jose",
	"luz",
	"thiago",
}

// namePrinter provides support for printing names
type namePrinter struct {
	name string
}

// Task implements the Worker interface
func (n namePrinter) Task() {
	// log.Println(n.name, id)
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	// create worker pool with 5 goroutines
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	// create a lot of goroutines competing to submit work to the pool
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				defer wg.Done()
				p.SendWork(np)
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
