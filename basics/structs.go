package basics

import "fmt"

// A struct is a collection of fields.

// Vertex is ...
type Vertex struct {
	X int
	Y int
}

// TStructs does ...
func TStructs() {
	fmt.Println("\n *** TStructs ***")
	v1 := Vertex{1, 2}
	fmt.Println(Vertex{1, 2})
	// Struct fields are accessed using a dot.
	v1.X = 4
	fmt.Println(v1.X)
	fmt.Println(v1)

	v2 := Vertex{1, 2}
	fmt.Println(v2)
	// Struct fields can be accessed through a struct pointer
	pointerV2 := &v2
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
	(*pointerV2).X = 1e4   // explicit dereference
	fmt.Println(v2)
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	pointerV2.X = 1e9
	fmt.Println(v2)
	
	var (
		// A struct literal denotes a newly allocated struct value
		v3 = Vertex{1, 2}  // has type Vertex
		// You can list just a subset of fields by using the Name: syntax. (In that case, the order of named fields is irrelevant.)
		v4 = Vertex{X: 1}  // v4.Y=0. Its value is implicit.
		v5 = Vertex{}      // v5.X=0 and v5.Y=0. Their values are implicit.
		// The special prefix & returns a pointer to the struct value
		pointerAny = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v3, pointerAny, v4, v5)

	fmt.Println()
	
}