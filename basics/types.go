package basics

import (
	"fmt"
	"math/cmplx"
)

const piFromTypes = 3.14

func zeroValuesVariables(){
	var i int
	var f float64
	var b bool
	var s string
	// Variables declared without an explicit initial value are given their zero value.
	// The zero value is:
	//   0                     for numeric types
	//   false                 for the boolean type
	//   "" (the empty string) for strings.
	fmt.Printf("%v %v %v ->%v<-\n", i, f, b, s)
}

// TTypes does
func TTypes() {
	fmt.Println("\n *** TTypes ***")

	// parenthesized, "factored" var statement.
	// Only need to use "var" 1 time
	var (
		v1              = false // default type: bool  . It will be true or false
		v2              = 5 // default type: int. It has 64 bits wide on 64-bit systems
		v3   int8       = 5
		v4   int16      = 5
		v5   int32      = 5
		aV5  rune       = 5 // alias for int32. Represents a Unicode code point
		v6   int64      = 5
		v7   uint       = 5 // 64 bits wide on 64-bit systems
		v8   uint8      = 5
		aV8  byte       = 5 // alias for uint8
		v9   uint16     = 5
		v10  uint32     = 5
		v11  uint64     = 5
		MaxInt uint64   = 1<<64 - 1
		v12  uintptr    = 44444 // 64 bits wide on 64-bit systems
		v13 float32     = 6.4
		v14             = 6.4 // default type: float64. The default value isn't float32!
		v15 complex64   = 5
		v16             = cmplx.Sqrt(-5 + 12i)  // Type: complex128
		v17             = "a text"   // Type: string
	)

	fmt.Printf("Types:\n" +
				"var1   bool       = %v Type:%T\n" +
				"var2   int        = %v Type:%T\n" +
				"var3   int8       = %v Type:%T\n" +
				"var4   int16      = %v Type:%T\n" +
				"var5   int32      = %v Type:%T\n" +
				"aVar5  rune       = %v Type:%T\n" +
				"var6   int64      = %v Type:%T\n" +
				"var7   uint       = %v Type:%T\n" +
				"var8   uint8      = %v Type:%T\n" +
				"aVar8  byte       = %v Type:%T\n" +
				"var9   uint16     = %v Type:%T\n" +
				"var10  uint32     = %v Type:%T\n" +
				"var11  uint64     = %v Type:%T\n" +
				"MaxInt uint64     = %v Type:%T\n" +
				"var12  uintptr    = %v Type:%T\n" +
				"var13 float32     = %v Type:%T\n" +
				"var14 float64     = %v Type:%T\n" +
				"var15 complex64   = %v Type:%T\n" +
				"var16 complex128  = %v Type:%T\n" +
				"var17 string      = %v Type:%T\n", v1,v1, v2,v2, v3,v3, v4,v4, v5,v5, aV5,aV5, v6,v6, v7,v7, v8,v8, aV8,aV8, v9,v9, v10,v10, v11,v11, MaxInt,MaxInt, v12,v12, v13,v13, v14,v14, v15,v15, v16,v16, v17,v17)
	zeroValuesVariables()
	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("const World = ", World)
	fmt.Println("outer const piFromTypes = ", piFromTypes)
	const Truth = true
	fmt.Println("const Truth = ", Truth)
	fmt.Println()
}