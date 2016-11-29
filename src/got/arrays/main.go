// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

func main() {

	// Declare an array of 5 strings set to its zero value.
	var arr1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	arr2 := [5]string{"Ed", "Russ", "Mike", "Kevin", "Pat"}

	// Assign the populated array to the array of zero values.
	arr1 = arr2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	/*
		// using range makes a copy of the string (pointer, length)
		for _, elem := range arr1 {
			fmt.Printf("%v : %p\n", elem, elem)
		}
	*/
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%v : %p\n", arr1[i], &arr1[i])
	}
}
