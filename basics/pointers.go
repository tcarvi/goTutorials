package basics

import (
	"fmt"
)

/*
 * A Pointer points to variables of a given type, called the base type of the pointer.
 * The value of an uninitialized pointer is nil.
 * A pointer holds the memory address of a value.
 * Go has no pointer arithmetic.
 */

// TPointers apresenta Pointers: its dataType, initiation, address operator e dereferencing.
func TPointers() {
	fmt.Println("\n*** TPointers ***")

	// Variables
	var v1 = 1 // Default infered type: int
	var v2 = 2 // Default infered type: int

	/*
	 * "var" identifier PointerType
	 *    PointerType = "*" BaseType
	 *    BaseType: Type
	 */
	var ptr1 *int // ptr1 is a pointer to an int value. It's value, without initiation, is nil.
	fmt.Println("Value pointed by ptr1 (pointer to v1), without initiation, is ", ptr1)

	/*
	 * & is the address operator ( pointer to its operand )
	 */
	ptr1 = &v1 // ptr1 now has the adress pointer to v1
	fmt.Println("Value of v1, pointed now by ptr1, is ", *ptr1)
	// For an operand v1 of type T, the address operation &v1 generates a pointer of type *T to v1.
	// The operand must be addressable, that is, either:
	//   a variable, pointer indirection,
	//   or slice indexing operation;
	//   or a field selector of an addressable struct operand;
	//   or an array indexing operation of an addressable array.
	// As an exception to the addressability requirement, v1 may also be a (possibly parenthesized) composite literal.
	// If the evaluation of v1 would cause a run-time panic, then the evaluation of &v1 does too.
	// For an operand v1 of pointer type *T, the pointer indirection *v1 denotes the variable of type T pointed to by v1.
	// If v1 is nil, an attempt to evaluate *v1 will cause a run-time panic.

	// "var" identifier = "&" identifier  // PointerType type is infered
	//    PointerType = "*" BaseType
	//    BaseType: Type
	// The variable whose pointerType is ("*" BaseType) is a pointer to a "Type" value.
	var ptr2 = &v2 // ptr2 is a pointer to an int value.
	fmt.Println("Value of v2, pointed by ptr2, is ", *ptr2)

	/*
	 *  *VarName is dereferencing operator ( VarName must be a pointer )
	 *   First type of dereferencing or "indirecting" operator. (To get value)
	 */
	v2 = *ptr1 // v2 now has the same value of v1
	fmt.Printf("New value of v2 is %v\n", v2)

	/*
	 * new(Type)
	 * The built-in function new takes a type T,
	 * allocates storage for a variable of that type at run time,
	 * and returns a value of type *T pointing to it. (default initial value)
	 */
	ptr3 := new(int)

	/*
	 *  *pointerIdentifier is a dereferencing operator ( pointerIdentifier must be a pointer )
	 *  Second type of dereferencing or "indirecting" operator. (To set a value)
	 */
	// There is a literal 3, not identified by a memory alocation.
	// pointer3 IS NOT nill (it has an address) and has the type *int
	// This "dereferencing operator" take the value 3 and copy its value to the address of pointer3.
	// This space in memory is set through the pointer pointer3
	*ptr3 = 3

	v4, v5, v6 := 4, 4, 12                        // := short assignment statement of tho int variables
	ptr4 := &v4                                   // := short assignment statement of a pointer to v4
	fmt.Println("Value pointed by ptr4 =", *ptr4) // Get value of a "dereferencing operator"
	var ptr5 *int = &v5                           // NOT inference of type *int
	*ptr5 = 5                                     // Set value of a "dereferencing operator"
	fmt.Println("Novo valor de v5 =", v5)

	var ptr6 *int = &v6 // NOT inference of type *int
	*ptr6 = *ptr6 / 2   // divide j through the pointer
	fmt.Println("Novo valor de v6 =", v6)

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
