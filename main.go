package main

import (
	"fmt"
)

/*
Function notes:
- if types match on params, only the last ought to be typed
- naked returns arent great for readability
*/
func addThree(n1 ,n2, n3 int) int {
	return n1 + n2 + n3
}

func returnMultiple() (string, string){
	return "hello", "world"
}

func nakedReturn() (x, y string) {
	x = "x string"
	y = "y string"
	return
}



func main() {

	// runSlices()
	// runArrays()
	// runPointers()
	// runMaps()
	// runStructs()
	// runVariables()
	// runAdditionalModules()

	// FUNCTIONS
	fmt.Println(addThree(23,10,22))
	fmt.Println(addThree(23,10,-7))
	string1, string2 := returnMultiple()
	fmt.Println("string1", string1, "string2: ", string2)
	string3, string4 := nakedReturn()
	fmt.Println(string3, string4)

	// FOR LOOPS

	for i:= 0; i < 3; i ++ {
		fmt.Print(i)
	}

	sum := 0
	for ; sum < 50; {
		sum += 10
		fmt.Print(sum)
	}

	// infinite loop
	// for {

	// }

	
}