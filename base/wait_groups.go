package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the goroutines launched to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	i := 0
	for i < 4 {
		wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		//
		go func() {
			worker(i)
			defer wg.Done()
		}()

		time.Sleep(1 * time.Millisecond) // sleep for at least 1 milli to allow each goroutine to instantiate
		i++
	}
	fmt.Printf("***** done******* i = %d\n", i)

	// Block until the WaitGroup counter goes back to 0 (when all the workers have notified they're done)
	wg.Wait()
}
