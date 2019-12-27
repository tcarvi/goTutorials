package basics // import "tcarvi.com/tutorial/basics"

import (
	"fmt"
)

// TForLoops do
func TForLoops(){
	fmt.Println("\n*** TForLoops ***")

	for i:=0; i<10; i++ {
		fmt.Printf("completeLoop%v\n", i)
	}

	fmt.Println()

	// a slice
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. 
	// The first is the index, 
	// and the second is a copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// You can use only the index, omitting  the returned value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// You can skip the index or value by assigning to _.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	// You can skip the index or value by assigning to _.
	for i, _ := range pow {
		fmt.Printf("%d\n", i)
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