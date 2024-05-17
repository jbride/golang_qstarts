package main

import (
	"fmt"
	"time"
)

/*  Overview
 *    Channels are the pipes that connect concurrent goroutines
 *    You can send values into channels from one goroutine and receive those values into another goroutine
 */

/* By default channels are unbuffered:
 *	 They will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
 * Buffered channels accept a limited number of values without a corresponding receiver for those values.
 */
func feedChannelBuffer(cBuffer chan string) {
	cBuffer <- "buffered"
	cBuffer <- "channel"
}

func main() {

	// Create a new channel w/ make(chan val-type).
	messages := make(chan string)
	go func() {
		messages <- "ping" // Send a value into a channel using the channel <- syntax
	}()

	// The <-channel syntax receives a value from the channel
	// NOTE: by default sends & receives block until both the sender and receiver are ready
	// Subsequently, our program can block for the "ping" message w/out having to use any other synchronization
	fmt.Printf("main() channel message from different thread = %s\n", <-messages)

	cBuffer := make(chan string, 2)
	go feedChannelBuffer(cBuffer)
	time.Sleep(5 * time.Second)
	fmt.Printf("cBuffers = %s , %s\n", <-cBuffer, <-cBuffer)
}
