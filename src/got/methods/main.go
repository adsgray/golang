// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p *player /* receiver */) average() float64 /* return type */ {
	// what about 0 at bats? that's not a player.
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{name: "Tim Raines", atBats: 9373, hits: 2834},
		{name: "Ken Griffy Jr.", atBats: 8373, hits: 2429},
		{name: "Andres Gallaraga", atBats: 9171, hits: 2313},
	}

	// Display the batting average for each player in the slice.
	for i := range players {
		fmt.Printf("Player: %v, average: %.3v\n", players[i].name, players[i].average())
	}

	fmt.Println("=========================")

	for _, p := range players {
		fmt.Printf("Player: %v, average: %.3v\n", p.name, p.average())
	}
}
