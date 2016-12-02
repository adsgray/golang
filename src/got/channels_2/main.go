// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
	//"sync"
	"time"
)

// Declare constant for number of goroutines .
const numroutines = 51

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffer channel with a buffer for
	// each goroutine to be created.
	c := make(chan int, numroutines)

	// Iterate and launch each goroutine.
	for ct := 0; ct < numroutines; ct++ {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func(num int) {
			fmt.Println(num, "sending")
			c <- rand.Intn(100)
		}(ct)

	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	var waitingfor int = numroutines

	// Iterate receiving each value until they are all received.
	for waitingfor > 0 {
		i := <-c
		fmt.Println("main received", i)
		waitingfor -= 1
	}
}
