// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"fmt"
	"sync"
)

// Declare a wait group variable.
var wg sync.WaitGroup

func main() {

	// Create an unbuffered channel.
	c := make(chan int)

	// Set the waitgroup, one for each goroutine.
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("A", c)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("B", c)
		wg.Done()
	}()

	// Send a value to start the counting.
	c <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, c chan int /* parameters */) {
	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.
		i, ok := <-c
		if !ok {
			fmt.Println(name, "channel closed, returning")
			return
		}

		// Display the value.
		fmt.Println(name, i)

		// Terminate when the value is 10.
		if i == 10 {
			fmt.Println(name, "i is 10, closing channel and returning")
			close(c)
			return
		}

		// Increment the value and send it
		// over the channel.
		i += 1
		c <- i
	}
}
