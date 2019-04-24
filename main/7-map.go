package main

import "fmt"

type Location struct {
	Lat, Long float64
}

var m map[string]Location

func main() {
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Location)
	m["Bell Labs"] = Location{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Can also define a map when declared
	var m2 = map[string]Location{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	// Delete a key/value in a map
	delete(m2, "Bell Labs")
	fmt.Println("The value:", m2["Bell Labs"])

	// Test if a key is exist, and also get the value
	v, ok := m2["Bell Labs"]
	fmt.Println("The value:", v, "Exist?", ok)
}
