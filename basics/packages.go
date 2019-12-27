// Every Go program is made up of packages.
// By convention, the package name is the same as the last element of the import path. 
// For instance, the "math/rand" package comprises files that begin with the statement package rand.

package basics // import "tcarvi.com/tutorial/basics"

// This program is using other packages
// It uses packages with import paths "fmt" and "math/rand".
import (
	"fmt"
	"math/rand"
)

// Programs start running in package main.
// It executes the code of "func main(){}"

// TPackages do
func TPackages(){
	fmt.Println("\n*** TPackages ***")
	fmt.Printf("Use of package rand, in directory math. Rand number = %v", rand.Intn(10))
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