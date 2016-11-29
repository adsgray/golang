// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	user1 := user{
		name:  "James",
		email: "james@jamesco.biz",
		age:   8,
	}

	// Display the field values.
	fmt.Println(user1.name)
	fmt.Println(user1.email)
	fmt.Println(user1.age)

	// Declare a variable using an anonymous struct.
	user2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Andrew",
		email: "andrew@andrewco.biz",
		age:   100,
	}

	// Display the field values.
	fmt.Println(user2.name)
	fmt.Println(user2.email)
	fmt.Println(user2.age)
}
