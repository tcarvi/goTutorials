package basics

import (
	"fmt"
	"strings"
)

// https://blog.golang.org/go-slices-usage-and-internals

// TSlices does
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
// In practice, slices are much more common than arrays.
func TSlices() {
	// An array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// The type []T is a slice with elements of type T.
	// A slice is formed by specifying two indices, a low and high+1 bound, separated by a colon: a[low : high+1]
	// The following expression creates a slice which includes elements 1 through 3 of the array primes:
	// A slice does not store any data, it just describes a section of an underlying array.
	// The first element of the array referenced by a slice is in the 0 index
	var s []int = primes[1:4]
	fmt.Println(s)

	// An array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// Two slices that reference the same array
	a := names[0:2]  // [John Paul]
	b := names[1:3]  // [Paul George]
	fmt.Println(a, b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same reference to the underlying array will see those changes.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// a slice that references the array [6]int{2, 3, 5, 7, 11, 13}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// a slice that references the array [6]bool{true, false, true, true, false, true}
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	// a slice that references the structs 
	// de tipo:
	//        {
	//	        i int 
	//          b bool 
	//        } 
	//  e de valores:
	//{ 
	//  {2, true},
	//  {3, false},
	//  {5, true},
	//  {7, true},
	//  {11, false},
	//  {13, true},
	//}
	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)
	// A chamada de elementos do slice é semelhante à chamada de elementos de array.
	fmt.Println(s1[0])

	// When slicing, you may omit the high or low bounds 
	// to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	s2 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s2)

	// Equivalent slices 
	// len() clanculates size of an array or of a slice
	var hightBoundPlus1 = len(s2)
	fmt.Println(s2[0:hightBoundPlus1])
	fmt.Println(s2[0:hightBoundPlus1])
	fmt.Println(s2[0:hightBoundPlus1])
	fmt.Println(s2[:hightBoundPlus1])
	fmt.Println(s2[0:])
	fmt.Println(s2[:])

	s2 = s2[1:4]
	fmt.Println(s2)

	s2 = s2[:2]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice("s3", s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice("s3", s3)

	// Extend its length.
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	s3 = s3[:4]
	printSlice("s3", s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice("s3", s3)

	var s4 []int
	// A nil slice has a length and capacity of 0 and has no underlying array.
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil {
		// The zero value of a slice is nil.
		fmt.Println("nil!")
		fmt.Printf("Type = %T\n", s4)
	}

	// Slices can be created with the built-in make function; 
	// this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	s5 := make([]int, 5)  //len(s5)=5
	printSlice("s5", s5)

	// To specify a capacity, pass a third argument to make:
	s6 := make([]int, 0, 5) // //len(s6)=0 cap(s6)=5
	printSlice("s6", s6)

	s7 := s6[:2]
	printSlice("s7", s7)

	s8 := s7[2:5]
	printSlice("s8", s8)

	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	s9 := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	s9[0][0] = "X"
	s9[2][2] = "O"
	s9[1][2] = "X"
	s9[1][0] = "O"
	s9[0][2] = "X"

	for i := 0; i < len(s9); i++ {
		fmt.Printf("%s\n", strings.Join(s9[i], " "))
	}

	var s10 []int
	printSlice("10", s10)

	// Go provides a built-in append function.
	// append works on nil slices.
	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, 
	// and the rest are T values to append to the slice.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	s10 = append(s10, 0)
	printSlice("s10", s10)

	// The slice grows as needed.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
	s10 = append(s10, 1)
	printSlice("s10", s10)

	// We can add more than one element at a time.
	s10 = append(s10, 2, 3, 4)
	printSlice("s10", s10)

}

func printSlice(name string, s []int) {
	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains. len(s)
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Printf("%v: len=%d cap=%d %v\n", name, len(s), cap(s), s)
}