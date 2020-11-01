package work

import (
	"sync"
)

// Worker must be implemented by types that want to use
// the worker pool.
type Worker interface {
	Task()
}

// Pool provides a set of goroutines that can execute any worker
// tasks that are submitted.
type Pool struct {
	workChan chan Worker
	wg       sync.WaitGroup
}

// New creates a new worker pool with a defined number of workers
func New(maxGoroutines int) *Pool {
	p := Pool{
		workChan: make(chan Worker),
	}

	// set number of goroutines to wait for
	// create 'maxGoroutines' number of goroutines, the goroutines
	// receive Worker interface values and calls the Task method on them.
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.workChan {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// SendWork sends work to channel
func (p *Pool) SendWork(w Worker) {
	p.workChan <- w
}

// Shutdown waits for all the goroutines to finish
func (p *Pool) Shutdown() {
	close(p.workChan)
	p.wg.Wait()
}
