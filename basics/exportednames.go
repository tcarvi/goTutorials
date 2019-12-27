package basics

import (
	"fmt"
	"math"
)

// TExportedNames does ...
func TExportedNames() {
	fmt.Println("\n *** TExportedNames ***")
	// A name is exported if it begins with a capital letter. 
	// Pi is exported from the math package.
	// You can't use math.pi in other package, but you can use math.Pi
	// Any "unexported" names are not accessible from outside the package.
	fmt.Printf("Uso of exported name Pi: math.Pi = %v", math.Pi)
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