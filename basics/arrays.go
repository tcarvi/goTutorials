package basics

import "fmt"

// TArrays resume https://golang.org/ref/spec#Array_types
func TArrays() {

	// Definição: - Sequência numerada
	//            - Apresenta quantidade definida de índices numéricos (ArrayLength)
	//            - Apresenta elementos de apenas 1 tipo (ElementType)
	// "Variable Declaration":
	// "var" idenfifier ArrayType
	//     identifier = letter { letter | unicode_digit }
	//     ArrayType = "[" ArrayLength "]" ElementType
	//     ArrayLength = Expression
	//     ElementType = Type
	// Depois de declarada, não se pode mudar o ArrayLength ou o ElementType de um Array
	// "Declaration" com "Initialization":
	// "var" idenfifier = [ArrayLength]ElementType
	// ArrayLength deve ser número "int" não-negativo
	// Deve-se usar a "built-in function" len para avaliar o ArrayLength
	// Os índices são iniados por zero e terminam em ( len(arrayIdentifier) - 1 )
	// A estrutura de dados do Array é unidimensional, mas pode ser organizada/composta para representar "multi-dimensional types"

	// Exemplos de "Variable Declaration":
	var arrayIdentifier01 [3]bool
	var arrayIdentifier02 [3]int //  int is the "integer" default type. int is 64 bits wide on 64-bit systems
	var arrayIdentifier03 [3]int8
	var arrayIdentifier04 [3]int16
	var arrayIdentifier05 [3]int32
	var arrayIdentifier06 [3]rune // alias para int32. Represents a Unicode code point.
	var arrayIdentifier07 [3]int64
	var arrayIdentifier08 [3]uint // uint is 64 bits wide on 64-bit systems
	var arrayIdentifier09 [3]uint16
	var arrayIdentifier10 [3]uint32
	var arrayIdentifier11 [3]uint64
	var arrayIdentifier12 [3]uintptr // uintptr is 64 bits wide on 64-bit systems
	var arrayIdentifier13 [3]float32
	var arrayIdentifier14 [3]float64 // float64 is the "float" default type. The default value isn't float32!
	//var arrayIdentifier15 [3]complex64
	//var arrayIdentifier16 [3]complex128
	var arrayIdentifier17 [3]string
	//var arrayIdentifier18 [3]struct{ fieldd1, field2 int }
	//var arrayIdentifier19 [3][4]string
	//var arrayIdentifier20 [3][4][5]string  // same as [2]([2]([2]string))

	// Exemplos de "Variable initialization":
	arrayIdentifier01 = [3]bool{true, false, true}
	arrayIdentifier02 = [3]int{0, 2, 4} //  int is the "integer" default type. int is 64 bits wide on 64-bit systems
	arrayIdentifier03 = [3]int8{0, 2, 4}
	arrayIdentifier04 = [3]int16{0, 2, 4}
	arrayIdentifier05 = [3]int32{0, 2, 4}
	arrayIdentifier06 = [3]rune{0, 2, 4} // alias para int32. Represents a Unicode code point.
	arrayIdentifier07 = [3]int64{0, 2, 4}
	arrayIdentifier08 = [3]uint{0, 2, 4} // uint is 64 bits wide on 64-bit systems
	arrayIdentifier09 = [3]uint16{0, 2, 4}
	arrayIdentifier10 = [3]uint32{0, 2, 4}
	arrayIdentifier11 = [3]uint64{0, 2, 4}
	arrayIdentifier12 = [3]uintptr{04567, 23456, 433445456} // uintptr is 64 bits wide on 64-bit systems
	arrayIdentifier13 = [3]float32{0.7, 2.7, 4.7}
	arrayIdentifier14 = [3]float64{0.7, 2.7, 4.7} // float64 is the "float" default type. The default value isn't float32!
	//arrayIdentifier15 = [3]complex64{true, false, true}
	//arrayIdentifier16 = [3]complex128{true, false, true}
	arrayIdentifier17 = [3]string{"str1", "str2", "str3"}
	/* arrayIdentifier18 = [3]struct{ fieldd1, field2 int }
		arrayIdentifier19 = [3][4]string{
											{{"Grupo1de3 indice1"}, {"Grupo1de3 indice2"}, {"Grupo1de3 indice3"}, {"Grupo1de3 indice4"}},
			  								{{"Grupo2de3 indice1"}, {"Grupo2de3 indice2"}, {"Grupo2de3 indice3"}, {"Grupo2de3 indice4"}},
											{{"Grupo3de3 indice1"}, {"Grupo3de3 indice2"}, {"Grupo3de3 indice3"}, {"Grupo3de3 indice4"}}
	   									}
		arrayIdentifier20 = [3][4][5]string{
												{
													{{"Grupo1 Subgrupo1de4 indice1"}, {"Grupo1 Subgrupo1de4 indice2"}, {"Grupo1 Subgrupo1de4 indice3"}, {"Grupo1 Subgrupo1de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}},
													{{"Grupo1 Subgrupo2de4 indice1"}, {"Grupo1 Subgrupo2de4 indice2"}, {"Grupo1 Subgrupo2de4 indice3"}, {"Grupo1 Subgrupo2de4 indice4"}, {"Grupo1 Subgrupo2de4 indice5"}},
													{{"Grupo1 Subgrupo3de4 indice1"}, {"Grupo1 Subgrupo3de4 indice2"}, {"Grupo1 Subgrupo3de4 indice3"}, {"Grupo1 Subgrupo3de4 indice4"}, {"Grupo1 Subgrupo3de4 indice5"}},
													{{"Grupo1 Subgrupo4de4 indice1"}, {"Grupo1 Subgrupo4de4 indice2"}, {"Grupo1 Subgrupo4de4 indice3"}, {"Grupo1 Subgrupo4de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}}
												},
												{
													{{"Grupo2 Subgrupo1de4 indice1"}, {"Grupo2 Subgrupo1de4 indice2"}, {"Grupo2 Subgrupo1de4 indice3"}, {"Grupo2 Subgrupo1de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}},
													{{"Grupo2 Subgrupo2de4 indice1"}, {"Grupo2 Subgrupo2de4 indice2"}, {"Grupo2 Subgrupo2de4 indice3"}, {"Grupo2 Subgrupo2de4 indice4"}, {"Grupo2 Subgrupo2de4 indice5"}},
													{{"Grupo2 Subgrupo3de4 indice1"}, {"Grupo2 Subgrupo3de4 indice2"}, {"Grupo2 Subgrupo3de4 indice3"}, {"Grupo2 Subgrupo3de4 indice4"}, {"Grupo2 Subgrupo3de4 indice5"}},
													{{"Grupo2 Subgrupo4de4 indice1"}, {"Grupo2 Subgrupo4de4 indice2"}, {"Grupo2 Subgrupo4de4 indice3"}, {"Grupo2 Subgrupo4de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}}
												},
												{
													{{"Grupo3 Subgrupo1de4 indice1"}, {"Grupo3 Subgrupo1de4 indice2"}, {"Grupo3 Subgrupo1de4 indice3"}, {"Grupo3 Subgrupo1de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}},
													{{"Grupo3 Subgrupo2de4 indice1"}, {"Grupo3 Subgrupo2de4 indice2"}, {"Grupo3 Subgrupo2de4 indice3"}, {"Grupo3 Subgrupo2de4 indice4"}, {"Grupo3 Subgrupo2de4 indice5"}},
													{{"Grupo3 Subgrupo3de4 indice1"}, {"Grupo3 Subgrupo3de4 indice2"}, {"Grupo3 Subgrupo3de4 indice3"}, {"Grupo3 Subgrupo3de4 indice4"}, {"Grupo3 Subgrupo3de4 indice5"}},
													{{"Grupo3 Subgrupo4de4 indice1"}, {"Grupo3 Subgrupo4de4 indice2"}, {"Grupo3 Subgrupo4de4 indice3"}, {"Grupo3 Subgrupo4de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}}
												}
										   } */

	// Exemplos de "Variable Declaration" com "Initialization", FORA a do escopo de funções:
	var arrayIdentifier21 = [3]bool{true, false, true}
	var arrayIdentifier22 = [3]int{0, 2, 4} //  int is the "integer" default type. int is 64 bits wide on 64-bit systems
	var arrayIdentifier23 = [3]int8{0, 2, 4}
	var arrayIdentifier24 = [3]int16{0, 2, 4}
	var arrayIdentifier25 = [3]int32{0, 2, 4}
	var arrayIdentifier26 = [3]rune{0, 2, 4} // alias para int32. Represents a Unicode code point.
	var arrayIdentifier27 = [3]int64{0, 2, 4}
	var arrayIdentifier28 = [3]uint{0, 2, 4} // uint is 64 bits wide on 64-bit systems
	var arrayIdentifier29 = [3]uint16{0, 2, 4}
	var arrayIdentifier30 = [3]uint32{0, 2, 4}
	var arrayIdentifier31 = [3]uint64{0, 2, 4}
	var arrayIdentifier32 = [3]uintptr{04567, 23456, 433445456} // uintptr is 64 bits wide on 64-bit systems
	var arrayIdentifier33 = [3]float32{0.7, 2.7, 4.7}
	var arrayIdentifier34 = [3]float64{0.7, 2.7, 4.7} // float64 is the "float" default type. The default value isn't float32!
	//var arrayIdentifier35 = [3]complex64{true, false, true}
	//var arrayIdentifier36 = [3]complex128{true, false, true}
	var arrayIdentifier37 = [3]string{"str1", "str2", "str3"}
	/* var arrayIdentifier38 = [3]struct{ fieldd1, field2 int }
	var arrayIdentifier39 = [3][4]string{
											{{"Grupo1de3 indice1"}, {"Grupo1de3 indice2"}, {"Grupo1de3 indice3"}, {"Grupo1de3 indice4"}},
											{{"Grupo2de3 indice1"}, {"Grupo2de3 indice2"}, {"Grupo2de3 indice3"}, {"Grupo2de3 indice4"}},
											{{"Grupo3de3 indice1"}, {"Grupo3de3 indice2"}, {"Grupo3de3 indice3"}, {"Grupo3de3 indice4"}}
										}
	var arrayIdentifier40 = [3][4][5]string{
												{
													{{"Grupo1 Subgrupo1de4 indice1"}, {"Grupo1 Subgrupo1de4 indice2"}, {"Grupo1 Subgrupo1de4 indice3"}, {"Grupo1 Subgrupo1de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}},
													{{"Grupo1 Subgrupo2de4 indice1"}, {"Grupo1 Subgrupo2de4 indice2"}, {"Grupo1 Subgrupo2de4 indice3"}, {"Grupo1 Subgrupo2de4 indice4"}, {"Grupo1 Subgrupo2de4 indice5"}},
													{{"Grupo1 Subgrupo3de4 indice1"}, {"Grupo1 Subgrupo3de4 indice2"}, {"Grupo1 Subgrupo3de4 indice3"}, {"Grupo1 Subgrupo3de4 indice4"}, {"Grupo1 Subgrupo3de4 indice5"}},
													{{"Grupo1 Subgrupo4de4 indice1"}, {"Grupo1 Subgrupo4de4 indice2"}, {"Grupo1 Subgrupo4de4 indice3"}, {"Grupo1 Subgrupo4de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}}
												},
												{
													{{"Grupo2 Subgrupo1de4 indice1"}, {"Grupo2 Subgrupo1de4 indice2"}, {"Grupo2 Subgrupo1de4 indice3"}, {"Grupo2 Subgrupo1de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}},
													{{"Grupo2 Subgrupo2de4 indice1"}, {"Grupo2 Subgrupo2de4 indice2"}, {"Grupo2 Subgrupo2de4 indice3"}, {"Grupo2 Subgrupo2de4 indice4"}, {"Grupo2 Subgrupo2de4 indice5"}},
													{{"Grupo2 Subgrupo3de4 indice1"}, {"Grupo2 Subgrupo3de4 indice2"}, {"Grupo2 Subgrupo3de4 indice3"}, {"Grupo2 Subgrupo3de4 indice4"}, {"Grupo2 Subgrupo3de4 indice5"}},
													{{"Grupo2 Subgrupo4de4 indice1"}, {"Grupo2 Subgrupo4de4 indice2"}, {"Grupo2 Subgrupo4de4 indice3"}, {"Grupo2 Subgrupo4de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}}
												},
												{
													{{"Grupo3 Subgrupo1de4 indice1"}, {"Grupo3 Subgrupo1de4 indice2"}, {"Grupo3 Subgrupo1de4 indice3"}, {"Grupo3 Subgrupo1de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}},
													{{"Grupo3 Subgrupo2de4 indice1"}, {"Grupo3 Subgrupo2de4 indice2"}, {"Grupo3 Subgrupo2de4 indice3"}, {"Grupo3 Subgrupo2de4 indice4"}, {"Grupo3 Subgrupo2de4 indice5"}},
													{{"Grupo3 Subgrupo3de4 indice1"}, {"Grupo3 Subgrupo3de4 indice2"}, {"Grupo3 Subgrupo3de4 indice3"}, {"Grupo3 Subgrupo3de4 indice4"}, {"Grupo3 Subgrupo3de4 indice5"}},
													{{"Grupo3 Subgrupo4de4 indice1"}, {"Grupo3 Subgrupo4de4 indice2"}, {"Grupo3 Subgrupo4de4 indice3"}, {"Grupo3 Subgrupo4de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}}
												}
											} */

	// Exemplos de "Variable Declaration" com "Initialization", DENTRO do escopo de funções:
	arrayIdentifier41 := [3]bool{true, false, true}
	arrayIdentifier42 := [3]int{0, 2, 4} //  int is the "integer" default type. int is 64 bits wide on 64-bit systems
	arrayIdentifier43 := [3]int8{0, 2, 4}
	arrayIdentifier44 := [3]int16{0, 2, 4}
	arrayIdentifier45 := [3]int32{0, 2, 4}
	arrayIdentifier46 := [3]rune{0, 2, 4} // alias para int32. Represents a Unicode code point.
	arrayIdentifier47 := [3]int64{0, 2, 4}
	arrayIdentifier48 := [3]uint{0, 2, 4} // uint is 64 bits wide on 64-bit systems
	arrayIdentifier49 := [3]uint16{0, 2, 4}
	arrayIdentifier50 := [3]uint32{0, 2, 4}
	arrayIdentifier51 := [3]uint64{0, 2, 4}
	arrayIdentifier52 := [3]uintptr{04567, 23456, 433445456} // uintptr is 64 bits wide on 64-bit systems
	arrayIdentifier53 := [3]float32{0.7, 2.7, 4.7}
	arrayIdentifier54 := [3]float64{0.7, 2.7, 4.7} // float64 is the "float" default type. The default value isn't float32!
	//arrayIdentifier55 := [3]complex64{true, false, true}
	//arrayIdentifier56 := [3]complex128{true, false, true}
	arrayIdentifier57 := [3]string{"str1", "str2", "str3"}
	/* arrayIdentifier58 := [3]struct{ fieldd1, field2 int }
	arrayIdentifier59 := [3][4]string{
								       {{"Grupo1de3 indice1"}, {"Grupo1de3 indice2"}, {"Grupo1de3 indice3"}, {"Grupo1de3 indice4"}},
								  	   {{"Grupo2de3 indice1"}, {"Grupo2de3 indice2"}, {"Grupo2de3 indice3"}, {"Grupo2de3 indice4"}},
								       {{"Grupo3de3 indice1"}, {"Grupo3de3 indice2"}, {"Grupo3de3 indice3"}, {"Grupo3de3 indice4"}}
								  }
	arrayIdentifier60 := [3][4][5]string{
										     {
											    {{"Grupo1 Subgrupo1de4 indice1"}, {"Grupo1 Subgrupo1de4 indice2"}, {"Grupo1 Subgrupo1de4 indice3"}, {"Grupo1 Subgrupo1de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}},
												{{"Grupo1 Subgrupo2de4 indice1"}, {"Grupo1 Subgrupo2de4 indice2"}, {"Grupo1 Subgrupo2de4 indice3"}, {"Grupo1 Subgrupo2de4 indice4"}, {"Grupo1 Subgrupo2de4 indice5"}},
												{{"Grupo1 Subgrupo3de4 indice1"}, {"Grupo1 Subgrupo3de4 indice2"}, {"Grupo1 Subgrupo3de4 indice3"}, {"Grupo1 Subgrupo3de4 indice4"}, {"Grupo1 Subgrupo3de4 indice5"}},
												{{"Grupo1 Subgrupo4de4 indice1"}, {"Grupo1 Subgrupo4de4 indice2"}, {"Grupo1 Subgrupo4de4 indice3"}, {"Grupo1 Subgrupo4de4 indice4"}, {"Grupo1 Subgrupo4de4 indice5"}}
											 },
											 {
												{{"Grupo2 Subgrupo1de4 indice1"}, {"Grupo2 Subgrupo1de4 indice2"}, {"Grupo2 Subgrupo1de4 indice3"}, {"Grupo2 Subgrupo1de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}},
												{{"Grupo2 Subgrupo2de4 indice1"}, {"Grupo2 Subgrupo2de4 indice2"}, {"Grupo2 Subgrupo2de4 indice3"}, {"Grupo2 Subgrupo2de4 indice4"}, {"Grupo2 Subgrupo2de4 indice5"}},
												{{"Grupo2 Subgrupo3de4 indice1"}, {"Grupo2 Subgrupo3de4 indice2"}, {"Grupo2 Subgrupo3de4 indice3"}, {"Grupo2 Subgrupo3de4 indice4"}, {"Grupo2 Subgrupo3de4 indice5"}},
												{{"Grupo2 Subgrupo4de4 indice1"}, {"Grupo2 Subgrupo4de4 indice2"}, {"Grupo2 Subgrupo4de4 indice3"}, {"Grupo2 Subgrupo4de4 indice4"}, {"Grupo2 Subgrupo4de4 indice5"}}
											 },
											 {
												{{"Grupo3 Subgrupo1de4 indice1"}, {"Grupo3 Subgrupo1de4 indice2"}, {"Grupo3 Subgrupo1de4 indice3"}, {"Grupo3 Subgrupo1de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}},
												{{"Grupo3 Subgrupo2de4 indice1"}, {"Grupo3 Subgrupo2de4 indice2"}, {"Grupo3 Subgrupo2de4 indice3"}, {"Grupo3 Subgrupo2de4 indice4"}, {"Grupo3 Subgrupo2de4 indice5"}},
												{{"Grupo3 Subgrupo3de4 indice1"}, {"Grupo3 Subgrupo3de4 indice2"}, {"Grupo3 Subgrupo3de4 indice3"}, {"Grupo3 Subgrupo3de4 indice4"}, {"Grupo3 Subgrupo3de4 indice5"}},
												{{"Grupo3 Subgrupo4de4 indice1"}, {"Grupo3 Subgrupo4de4 indice2"}, {"Grupo3 Subgrupo4de4 indice3"}, {"Grupo3 Subgrupo4de4 indice4"}, {"Grupo3 Subgrupo4de4 indice5"}}
											 }
										} */

	// Não se deve passar arrays para functions porque ocorrerá a cópia dos arrays
	// Deve-se passar slices destes arrays, como argumento de funções

	var sliceOfArray01 []bool = arrayIdentifier01[:]
	printArrayBool("arrayIdentifier01", sliceOfArray01)

	var sliceOfArray02 []int = arrayIdentifier02[:]
	printArrayInt("arrayIdentifier02", sliceOfArray02)

	var sliceOfArray03 []int8 = arrayIdentifier03[:]
	printArrayInt8("arrayIdentifier03", sliceOfArray03)

	var sliceOfArray04 []int16 = arrayIdentifier04[:]
	printArrayInt16("arrayIdentifier04", sliceOfArray04)

	var sliceOfArray05 []int32 = arrayIdentifier05[:]
	printArrayInt32("arrayIdentifier05", sliceOfArray05)

	var sliceOfArray06 []rune = arrayIdentifier06[:]
	printArrayRune("arrayIdentifier06", sliceOfArray06)

	var sliceOfArray07 []int64 = arrayIdentifier07[:]
	printArrayInt64("arrayIdentifier07", sliceOfArray07)

	var sliceOfArray08 []uint = arrayIdentifier08[:]
	printArrayUint("arrayIdentifier08", sliceOfArray08)

	var sliceOfArray09 []uint16 = arrayIdentifier09[:]
	printArrayUint16("arrayIdentifier09", sliceOfArray09)

	var sliceOfArray10 []uint32 = arrayIdentifier10[:]
	printArrayUint32("arrayIdentifier10", sliceOfArray10)

	var sliceOfArray11 []uint64 = arrayIdentifier11[:]
	printArrayUint64("arrayIdentifier11", sliceOfArray11)

	var sliceOfArray12 []uintptr = arrayIdentifier12[:]
	printArrayUintptr("arrayIdentifier12", sliceOfArray12)

	var sliceOfArray13 []float32 = arrayIdentifier13[:]
	printArrayFloat32("arrayIdentifier13", sliceOfArray13)

	var sliceOfArray14 []float64 = arrayIdentifier14[:]
	printArrayFloat64("arrayIdentifier14", sliceOfArray14)

	//var sliceOfArray15 []complex64 = arrayIdentifier15[:]
	//printArrayComplex64("arrayIdentifier15", sliceOfArray15)

	//var sliceOfArray16 []complex128= arrayIdentifier16[:]
	//printArrayComplex128("arrayIdentifier16", sliceOfArray16)

	var sliceOfArray17 []string = arrayIdentifier17[:]
	printArrayString("arrayIdentifier17", sliceOfArray17)

	// TODO Print de Struct e de multi-Arrays

	var sliceOfArray21 []bool = arrayIdentifier21[:]
	printArrayBool("arrayIdentifier21", sliceOfArray21)

	var sliceOfArray22 []int = arrayIdentifier22[:]
	printArrayInt("arrayIdentifier22", sliceOfArray22)

	var sliceOfArray23 []int8 = arrayIdentifier23[:]
	printArrayInt8("arrayIdentifier23", sliceOfArray23)

	var sliceOfArray24 []int16 = arrayIdentifier24[:]
	printArrayInt16("arrayIdentifier24", sliceOfArray24)

	var sliceOfArray25 []int32 = arrayIdentifier25[:]
	printArrayInt32("arrayIdentifier25", sliceOfArray25)

	var sliceOfArray26 []rune = arrayIdentifier26[:]
	printArrayRune("arrayIdentifier26", sliceOfArray26)

	var sliceOfArray27 []int64 = arrayIdentifier27[:]
	printArrayInt64("arrayIdentifier27", sliceOfArray27)

	var sliceOfArray28 []uint = arrayIdentifier28[:]
	printArrayUint("arrayIdentifier28", sliceOfArray28)

	var sliceOfArray29 []uint16 = arrayIdentifier29[:]
	printArrayUint16("arrayIdentifier29", sliceOfArray29)

	var sliceOfArray30 []uint32 = arrayIdentifier30[:]
	printArrayUint32("arrayIdentifier30", sliceOfArray30)

	var sliceOfArray31 []uint64 = arrayIdentifier31[:]
	printArrayUint64("arrayIdentifier31", sliceOfArray31)

	var sliceOfArray32 []uintptr = arrayIdentifier32[:]
	printArrayUintptr("arrayIdentifier32", sliceOfArray32)

	var sliceOfArray33 []float32 = arrayIdentifier33[:]
	printArrayFloat32("arrayIdentifier33", sliceOfArray33)

	var sliceOfArray34 []float64 = arrayIdentifier34[:]
	printArrayFloat64("arrayIdentifier34", sliceOfArray34)

	//var sliceOfArray35 []complex64= arrayIdentifier35[:]
	//printArrayComplex64("arrayIdentifier35", sliceOfArray35)

	//var sliceOfArray36 []complex128 = arrayIdentifier36[:]
	//printArrayComplex128("arrayIdentifier36", sliceOfArray36)

	var sliceOfArray37 []string = arrayIdentifier37[:]
	printArrayString("arrayIdentifier37", sliceOfArray37)

	// TODO Print de Struct e de multi-Arrays

	var sliceOfArray41 []bool = arrayIdentifier41[:]
	printArrayBool("arrayIdentifier41", sliceOfArray41)

	var sliceOfArray42 []int = arrayIdentifier42[:]
	printArrayInt("arrayIdentifier42", sliceOfArray42)

	var sliceOfArray43 []int8 = arrayIdentifier43[:]
	printArrayInt8("arrayIdentifier43", sliceOfArray43)

	var sliceOfArray44 []int16 = arrayIdentifier44[:]
	printArrayInt16("arrayIdentifier44", sliceOfArray44)

	var sliceOfArray45 []int32 = arrayIdentifier45[:]
	printArrayInt32("arrayIdentifier45", sliceOfArray45)

	var sliceOfArray46 []rune = arrayIdentifier46[:]
	printArrayRune("arrayIdentifier46", sliceOfArray46)

	var sliceOfArray47 []int64 = arrayIdentifier47[:]
	printArrayInt64("arrayIdentifier47", sliceOfArray47)

	var sliceOfArray48 []uint = arrayIdentifier48[:]
	printArrayUint("arrayIdentifier48", sliceOfArray48)

	var sliceOfArray49 []uint16 = arrayIdentifier49[:]
	printArrayUint16("arrayIdentifier49", sliceOfArray49)

	var sliceOfArray50 []uint32 = arrayIdentifier50[:]
	printArrayUint32("arrayIdentifier50", sliceOfArray50)

	var sliceOfArray51 []uint64 = arrayIdentifier51[:]
	printArrayUint64("arrayIdentifier51", sliceOfArray51)

	var sliceOfArray52 []uintptr = arrayIdentifier52[:]
	printArrayUintptr("arrayIdentifier52", sliceOfArray52)

	var sliceOfArray53 []float32 = arrayIdentifier53[:]
	printArrayFloat32("arrayIdentifier53", sliceOfArray53)

	var sliceOfArray54 []float64 = arrayIdentifier54[:]
	printArrayFloat64("arrayIdentifier54", sliceOfArray54)

	//var sliceOfArray55 []complex64= arrayIdentifier55[:]
	//printArrayComplex64("arrayIdentifier55", sliceOfArray55)

	//var sliceOfArray56 []complex128 = arrayIdentifier56[:]
	//printArrayComplex128("arrayIdentifier56", sliceOfArray56)

	var sliceOfArray57 []string = arrayIdentifier57[:]
	printArrayString("arrayIdentifier57", sliceOfArray57)

	// TODO Print de Struct e de multi-Arrays
}

func printArrayBool(name string, sliceOfArray []bool) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayInt(name string, sliceOfArray []int) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayInt8(name string, sliceOfArray []int8) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayInt16(name string, sliceOfArray []int16) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayInt32(name string, sliceOfArray []int32) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayRune(name string, sliceOfArray []rune) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayInt64(name string, sliceOfArray []int64) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayUint(name string, sliceOfArray []uint) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayUint16(name string, sliceOfArray []uint16) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayUint32(name string, sliceOfArray []uint32) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayUint64(name string, sliceOfArray []uint64) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayUintptr(name string, sliceOfArray []uintptr) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayFloat32(name string, sliceOfArray []float32) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayFloat64(name string, sliceOfArray []float64) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayComplex64(name string, sliceOfArray []complex64) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayComplex128(name string, sliceOfArray []complex128) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayString(name string, sliceOfArray []string) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayStruct(name string, sliceOfArray []struct{ fieldd1, field2 int }) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayDoubleArrayString(name string, sliceOfArray [3][4]string) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}

func printArrayTripleArrayString(name string, sliceOfArray [3][4][5]string) {
	fmt.Println("Elementos do array", name, ":")
	for i := 0; i < len(sliceOfArray); i++ {
		fmt.Println(sliceOfArray[i])
	}
}
