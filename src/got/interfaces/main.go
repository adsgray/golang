// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.
import "fmt"

// Declare the speaker interface with a single method called speak.
type speaker interface {
	speak()
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (english) speak() {
	fmt.Println("Hello World")
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a pointer receiver. "你好世界"
func (*chinese) speak() {
	fmt.Println("你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker /* Declare parameter */) {
	// Call the speak method from the speaker parameter.
	s.speak()
}

func main() {

	// Declare a variable of the interface speaker type
	// set to its zero value.
	var s1 speaker

	// Declare a variable of type english.
	var e1 english

	// Assign the english value to the speaker variable.
	s1 = e1

	// Call the speak method against the speaker variable.
	s1.speak()

	// Declare a variable of type chinese.
	var c1 chinese

	// Assign the chinese pointer to the speaker variable.
	//s1 = c1
	s1 = &c1

	// Call the speak method against the speaker variable.
	s1.speak()

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	{
		var e1 english
		e2 := &e1
		sayHello(e1)
		sayHello(e2)

		var c1 chinese
		c2 := &c1
		//sayHello(c1) pointer receiver so does not implement interface
		sayHello(&c1)
		sayHello(c2)
	}

}
