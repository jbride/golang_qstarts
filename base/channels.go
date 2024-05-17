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

const messageCount = 5

func feedChannelBuffer(cBuffer chan string) {
	cBuffer <- "buffered"
	cBuffer <- "channel"
}

func main() {

	// Create a new channel w/ make(chan val-type).
	messages := make(chan string)
	go func() {
		for i := 0; i < messageCount; i++ {
			message := fmt.Sprint("ping", i)
			messages <- message // Send a value into a channel; this will block until receiver is ready
			fmt.Printf("anonFunc() just sent message to channel:  %s\n", message)
		}
	}()

	// The <-channel syntax receives a value from the channel
	// NOTE: by default sends & receives block until both the sender and receiver are ready
	// Subsequently, this program can block for the "ping" message w/out having to use any other synchronization
	for i := 0; i < messageCount; i++ {
		fmt.Printf("main() channel message from different thread = %s\n", <-messages)
		time.Sleep(1 * time.Second)
	}

	// Channel Buffers
	cBuffer := make(chan string, 2)
	go feedChannelBuffer(cBuffer)
	time.Sleep(3 * time.Second)
	fmt.Printf("\n\ncBuffers = %s , %s\n", <-cBuffer, <-cBuffer)
	//fmt.Printf("cBuffers = %s , %s . %s\n", <-cBuffer, <-cBuffer, <-cBuffer) //fatal runtime error: all goroutines are asleep - deadlock!
}
