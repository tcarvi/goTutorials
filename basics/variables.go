package basics

import (
	"fmt"
)

// The var statement declares a list of variables
// As in function argument lists, the type is last.
// A var statement can be at package or function level.
var v1, v2 bool

// A var declaration can include initializers, one per variable.
var v3, v4, v5 = true, false, "v5!"

// You CAN'T use a "short assignment statement (:=)" outside a function, 

// TVariables does ...
func TVariables() {
	fmt.Println("\n *** TVariables ***")

	// A var statement can be at package or function level.
	var v6 int

	// := short assignment statement
	// Used only inside a function, 
	// It can be used in place of a var declaration with implicit type.
	v7 := 7
	v8, v9, v10 := true, false, "v10!"

	// If an initializer is present, the type can be omitted; 
	// the variable will take the type of the initializer.
	var v11 = 11
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v, v6=%v, v7=%v, v8=%v, v9=%v, v10=%v, v11=%v", v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11)

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