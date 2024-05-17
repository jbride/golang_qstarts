package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(threadNum int) {
			// Have each thread will sleep for a random amount of time
			sleepTime := rand.Int31n(500)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			wg.Done()
			fmt.Printf("anonFunc() just finished thread:\t %d \t: sleep duration = %d\n", threadNum, sleepTime)
		}(i)
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
