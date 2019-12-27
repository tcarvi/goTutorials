package basics

import "fmt"

// Vertices does ...
type Vertices struct {
	Lat, Long float64
}
// A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
var m map[string]Vertices

// TMaps does ...
func TMaps() {
	// The make function returns a map of the given type, initialized and ready for use
	m = make(map[string]Vertices)
	m["Bell Labs"] = Vertices{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m2["Bell Labs"])
	fmt.Println(m2)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	m4 := make(map[string]int)

	// Insert or update an element in map m:
	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	// // Insert or update an element in map m:
	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	// Delete an element:
	delete(m4, "Answer")
	// If key is not in the map, then elem is the zero value for the map's element type.
	fmt.Println("The value:", m4["Answer"])

	// Retrieve an element:
	value =  m4["Answer"]

	// If key is in m, ok is true. If not, ok is false.
	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
