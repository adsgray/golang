// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var int1 []int

	// Appends numbers to the slice.
	int1 = append(int1, 1)
	int1 = append(int1, 2)
	int1 = append(int1, 3)
	int1 = append(int1, 4)

	// Display each value in the slice.
	for _, val := range int1 {
		fmt.Println(val)
	}

	fmt.Println("=========================")

	// Declare a slice of strings and populate the slice with names.
	s1 := []string{"Jen", "David", "James", "Alison", "Pat", "Ron", "Sandy"}

	// Display each index position and slice value.
	for i := 0; i < len(s1); i++ {
		fmt.Printf("index: %v value:%v\n", i, s1[i])
	}

	fmt.Println("=========================")

	// Take a slice of index 1 and 2 of the slice of strings.
	s2 := s1[1 : 1+2] // start plus length

	// Display each index position and slice values for the new slice.
	for i := 0; i < len(s2); i++ {
		fmt.Printf("index: %v value:%v\n", i, s2[i])
	}
}
