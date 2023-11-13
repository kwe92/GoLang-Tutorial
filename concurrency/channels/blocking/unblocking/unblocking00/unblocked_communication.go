package main

import "fmt"

func main() {

	// create channel of integers
	unbuffered := make(chan int)

	// launch a goroutine to write to the channel
	go func() {
		unbuffered <- 42

		// close the channel | will not deadlock if ranged over
		close(unbuffered)

	}()

	// in a separate goroutine read from the channel | main is technically the first goroutine launched

	for ele := range unbuffered {
		fmt.Println(ele)
	}

}

// Unbuffered Channels: Sender and Reciever Required

//   - you need at least two goroutines for a channel to communicate
//   - a sender and a receiver
//   - the two goroutines are synchronized to communicate at a certain point
