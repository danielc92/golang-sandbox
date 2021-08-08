package main

import "fmt"

func runPointers() {
	// POINTERS
	// Understanding * and & operands
	// * reveals a pointers real value
	// logging the pointer alone, will only show the memory address
	// & operand can be used to create a new pointer to a variable

	original := 10
	pointer := &original
	fmt.Println("Pointer address: ", pointer)
	fmt.Println("Pointers value (original): ", *pointer)
	// setting original through pointer
	*pointer = 11
	fmt.Println("Value should be 11: ", original)
}