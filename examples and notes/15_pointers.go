package main

// a pointer is a variable that stores the memory
// address of another variable. they are useful when
// you want to change a value in a function or you
// want to avoid making copy of large variables.

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// *int means it takes an integer pointer
func zeroptr(iptr *int) {
	// function dereferences the pointer from its memory
	// address to the current value at that address. Assigning
	// a value to a dereferenced pointer changes the value
	// at the referenced address.
	*iptr = 0
}

func main() {
	a := 10
	b := &a // b is a pointer to the memory address of a

	fmt.Println("a:", a)
	fmt.Println("b (memory address)", b) // memory address
	fmt.Println("b:", *b)                // accessing value

	*b = 11 // changing the value at the memory address of a
	fmt.Println("a:", a)

	// declaring i
	i := 1
	zeroval(i)
	fmt.Println("i:", i)

	// changing i by derefrecning its value at its memory address
	zeroptr(&i)
	fmt.Println("i after dereferencing its value at the memory address is", i)
}
