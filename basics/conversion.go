package basics

import (
	"fmt"
	"math"
)

// TConversion does ...
func TConversion() {
	fmt.Println("\n *** TConversion ***")
	var v1, v2 int = 3, 4
	// The expression T(v) converts the value v to the type T.
	// T=float64 , v=x*x + y*y   ; float64(x*x + y*y) gera type float64 
	var v3 = math.Sqrt(float64(v1*v1 + v2*v2))  // type: float64
	// T=uint , v=f   ; uint(f) gera type uint 
	var v4 = uint(v3)  // type: uint
	fmt.Println(v3, v4)
	v5 := 42
	v6 := float64(v5)
	v7 := uint(v6)
	fmt.Println(v6, v7)
}