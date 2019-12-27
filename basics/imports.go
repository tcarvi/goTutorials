package basics // import "tcarvi.com/tutorial/basics"

// parenthesized, "factored" import statement.
import (
	"fmt"
	"math"
)

// TImports do...
func TImports() {
	fmt.Println("\n *** TImports ***")
	fmt.Printf("Import of package 'math'. math.Sqrt(7) = %v", math.Sqrt(7))
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