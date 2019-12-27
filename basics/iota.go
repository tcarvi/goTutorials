package basics // import "tcarvi.com/tutorial/basics"

import (
	"fmt"
)

// TIota do
func TIota(){
	fmt.Println("\n*** TIota ***")

	type Grades int
	// We only need to define the first, which will be 0
	// The others will have the later added by 1
	const (
		A Grades = iota
		B
		C
		D
		F
	)
	fmt.Printf("A=%v B=%v C=%v D=%v F=%v", A, B, C, D, F)
	fmt.Println()
}

// The default format for %v, inside a fmt.Printf parameter, is:
// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p