package basics

import (
	"fmt"
	"math"
)

// add adiciona números
// function receiver: sem aglutinador de funções
// function parameters: x de type int            ( A function can take zero or more arguments )
//                     y de type int
// Return values: not named, de type int
func add(x int, y int) int {
	return x + y
}

// add adiciona números
// function receiver: sem aglutinador de funções
// function parameters: x de type int
//                      y de type int
// Return values: not named, de type int
func addWithSameTypeParameters(x, y int) int {
	return x + y
}

// swap troca posição e retorna valores
// function receiver: sem aglutinardor
// function parameters: x de type string
//                      y de type string
// Return values: not named, de type string           ( A function can return any number of results )
func swap(x, y string) (string, string) {
	return y, x
}

// split processa valores internamente e os retorna
// function receiver: sem aglutinador de funções
// function parameters: sum de type int
// Return values: x de type int       ( return values may be named)
//                y de type int       ( If so, they are treated as variables defined at the top of the function. )
func processAndReturn(i int) (x, y int) {
	x = i * 10 / 2
	y = i - x
	// A return statement without arguments returns the named return values. 
	// This is known as a "naked" return.
	return
}

// fn is passed as a parameter
// fn is a function value 
// fn has type: func(float64, float64) float64
// Function compute return a value processed internally by fn
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// the adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// TFunction does ...
func TFunction() {
	fmt.Println("\n *** TFunction ***")
	fmt.Printf("Use of internal function: add(42, 13) = %v\n", add(42, 13))
	fmt.Printf("Use of internal function: addWithSameTypeParameters(42, 13) = %v\n", addWithSameTypeParameters(42, 13))
	a, b := swap("42", "13")
	fmt.Printf("Use of internal function: swap('42', '13') = %v %v\n", a, b)
	c, d := processAndReturn(100)
	fmt.Printf("Use of internal function: processAndReturn(100): (x = i * 10 / 2 = %v) and (y = i - x = %v)", c, d)
	fmt.Println()

	// Function value. Variable of type func(x, y float64) float64
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// The default format for %v, inside a fmt.Printf parameter, is:
// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p