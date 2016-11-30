// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the match method for hockey, call into the match method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// match checks the value for the specified term.
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	goaliename string
}

// match checks the value for the specified term.
func (h hockey /* receiver type */) match(searchTerm string) bool {

	// Make sure you call into match method for the embedded sport type.
	if h.sport.match(searchTerm) {
		return true
	}

	// Implement the search for the new fields.
	return strings.Contains(h.goaliename, searchTerm)
}

func main() {

	// Define the term to match.
	m := "Plante"

	// Create a slice of matcher values and assign values
	// of the concrete hockey type.
	teams := []matcher{
		hockey{sport{"Leafs", "Toronto"}, "Bauer"},
		hockey{sport{"Canadiens", "Montreal"}, "Plante"},
		hockey{sport{"Black Hawks", "Chicago"}, "Hall"},
	}

	// Display what we are trying to match.
	fmt.Printf("looking for: %q\n", m)

	// Range of each matcher value and check the term.
	for _, t := range teams {
		if t.match(m) {
			fmt.Printf("%v matches\n", t)
		}
	}
}
