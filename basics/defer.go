package basics

import "fmt"

// TDefer does ...
func TDefer() {
	fmt.Println("\n *** TDefer ***")

	// A defer statement defers the execution of the inner call.
	// Only when the the surrounding function returns, the inner call start its execution.
	// Deferred function calls are pushed onto a stack. 
	// When the surrounding function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	defer fmt.Println() // Última chamada a ser executada: last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Execução antes das chamadas indicadas por defer")
}