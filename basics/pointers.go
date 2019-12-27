package basics // import "tcarvi.com/tutorial/basics"

import (
	"fmt"
)

// TPointers do
// A pointer holds the memory address of a value.
// Go has no pointer arithmetic.
func TPointers(){
	fmt.Println("\n*** TPointers ***")

	var v1 = 1 // Default infered type: int
	var v2 int
	// The type *T is a pointer to a T value. Its zero value is nil.
	var pointerV *int   // ip is a pointer to int. It's value, without initiation, is nil.
	fmt.Println("ip value, without initiation, is = ", pointerV)
	/*
	 * & is the address operator ( pointer to its operand )
	 */
	 pointerV = &v1       // pointerV now has the adress pointer to v1

	/*
	 *  *VarName is dereferencing operator ( VarName must be a pointer )
	 *   First type of dereferencing operator. (To get value)
	 */
	 v2 = *pointerV       // v2 is now 1

	fmt.Printf("v2 = %v\n", v2)

	/*
	 *  new(type)
	 *  Create a variable with default initial value
	 *  Then it return a pointer to this variable
	 */ 
	 ptr := new(int)
	 // Second type: "dereferencing" or "indirecting" operator. (To set a value)
	 // A space in memory is set through the pointer ptr
	 *ptr = 3

	 i, j := 42, 2701
	 p := &i         // point to i
	 fmt.Println(*p) // read i through the pointer
	 *p = 21         // set i through the pointer
	 fmt.Println(i)  // see the new value of i
 
	 p = &j         // point to j
	 *p = *p / 37   // divide j through the pointer
	 fmt.Println(j) // see the new value of j

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