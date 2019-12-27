package basics

import "fmt"

// Go has only one looping construct, the for loop.

// TLoop does ...
// The basic for loop has three components separated by semicolons:
//   the init statement: executed before the first iteration
//   the condition expression: evaluated before every iteration
//   the post statement: executed at the end of every iteration
// The init statement will often be a short variable declaration
// The variables declared there are visible only in the scope of the for statement.
// The loop will stop iterating once the boolean condition evaluates to false.
func TLoop() {
	fmt.Println("\n*** TLoop ***")
	fmt.Println("\n*** Loop 1 ***")
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("\n*** Loop 2 ***")
	sum2 := 1
	// The init and post statements are optional. 
	// Their space can be empty. Or you can drop the semicolons
	for ; sum2 < 20; {
		fmt.Println(sum2)
		sum2 += sum2
	}
	fmt.Println(sum2)
	fmt.Println("\n*** Loop 3 ***")
	sum3 := 1
	// Example of dropped the semicolons
	for sum3 < 20 {
		sum3 += sum3
	}
	fmt.Println(sum3)
	// Example of a infinite loop
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	// for {
	// }
}
