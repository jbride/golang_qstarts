package main

import (
	"fmt"
	"sync"
	"time"
)

const useWaitGroup = true
const mainThreadBlock = true

func workFunction1(id int) {
	fmt.Printf("workFunction1 %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("workFunction1 %d done\n", id)
}

func main() {

	// Without use of a WaitGroup, the main thread will terminate before any of the child threads actually execute
	for y := 0; y < 5; y++ {

		// Call an anonymous function on new thread
		go func() {
			fmt.Printf("naive thread %d\n", y)
		}()
	}

	if useWaitGroup {
		// This WaitGroup is used to wait for all the goroutines launched to finish.
		var wg sync.WaitGroup

		// Launch several goroutines and increment the WaitGroup counter for each
		i := 0
		for i < 4 {
			wg.Add(1)

			// Wrap the worker call in a closure that defers notification to WaitGroup until after completion of workFunction
			go func() {

				// A defer statement defers execution of a function until the surrounding function returns
				// The deferred call's args are evaluated immediately, but the function is not executed until the surrounding function returns
				defer wg.Done()
				workFunction1(i)
			}()

			// sleep for at least 1 milli to allow each goroutine to instantiate on a new thread prior to loop continuing on in primary thread
			time.Sleep(1 * time.Millisecond)
			i++
		}

		// Block until the WaitGroup counter goes back to 0 (when all the workers have notified they're done)
		// Notice the behavior of this quickstart when the main thread doesn't block
		if mainThreadBlock {
			wg.Wait()
		}
	}
	fmt.Printf("useWaitGroup = %t ; mainThreadBlock = %t\n", useWaitGroup, mainThreadBlock)
}
