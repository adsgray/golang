// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
func funcName(userin *user, newname string /* add pointer parameter, add value parameter */) {

	// Use the pointer to change the value that the
	// pointer points to.
	userin.name = newname
}

func main() {

	// Create a variable of type user and initialize each field.
	user1 := user{
		name:  "James",
		email: "james@jamesco.biz",
		age:   8,
	}

	// Display the value of the variable.
	fmt.Println(user1.name)
	fmt.Println(user1.email)
	fmt.Println(user1.age)

	// Share the variable with the function you declared above.
	funcName(&user1, "Patrick")

	// Display the value of the variable.
	fmt.Println(user1.name)
	fmt.Println(user1.email)
	fmt.Println(user1.age)
}
