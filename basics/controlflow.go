package basics

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	// the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// the variable declared in the if short statement are also available here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	// in the end of the braces, the variable v is out of scope
	return lim
}

// TControlFlow does ...
func TControlFlow() {
	fmt.Println("\n *** TControlFlow ***")
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Switch statement
	// It only runs the selected case, not all the cases that follow. 
	// The break statement that is needed at the end of each case in those languages is provided automatically in Go. 
	// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(time.Wednesday)
	switch time.Wednesday {
	case today + 0:
		// Nothing will be done. Switch finishes
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println()
}