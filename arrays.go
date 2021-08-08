package main

import "fmt"

func runArrays() {
	// ARRAYS
	// arrays are fixed size [n]T

	var flags [3]bool
	var names [2]string
	names = [2]string{"Daniel", "Corcoran"}
	fmt.Println(names)
	fmt.Println("Array length: ", len(flags))
	// setting array at particular index
	flags[0] = true
	flags[1] = false
	flags[2] = true
	fmt.Println("Array values: ", flags)

	// looping through arrays
	for index, value := range flags {
		fmt.Println("Looping through Array: ", index, value)
	}
}