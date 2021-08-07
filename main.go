package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

	// TIME MODULE
	fmt.Println("Get datetime now", time.Now())
	fmt.Println("Get unix now", time.Now().Unix()) // Note this is in seconds, not milliseconds
	fmt.Println("Get month now", time.Now().Month())
	fmt.Println("Get year now", time.Now().Year())

	// MATH MODULE
	fmt.Println("Finding the ceiling (rounding up)", math.Ceil(3.44))
	fmt.Println("Finding the floor (rounding down)", math.Floor(4.99))
	fmt.Println("Finding the square root", math.Sqrt(3.34))
	fmt.Println("Finding the power of 4^3", math.Pow(4, 3))
	fmt.Println("Math constants", math.Pi, math.SqrtE, math.E)
	danielsRandomNo := rand.Intn(20)
	fmt.Println("Generating random numbers", danielsRandomNo)

	// FUNCTIONS
	fmt.Println(addThree(23,10,22))
	fmt.Println(addThree(23,10,-7))
	string1, string2 := returnMultiple()
	fmt.Println("string1", string1, "string2: ", string2)
	string3, string4 := nakedReturn()
	fmt.Println(string3, string4)
	
	// VARIABLES

	// declare multiple
	var f,g,h int

	// declare only
	var b bool
	var x int32
	var s4 string

	// declare AND set
	var b2 bool = true

	// declare and set multiple
	var weight, height int32 = 80, 160

	// allow golang to infer type
	b3 := true
	s1 := "daniel"
	x1 := -2424356

	// casting float to int
	// be careful with losing data/information
	x2 := 2.22
	x3 := int(x2)

	/*
	 Zero values
	- 0 for numeric types (int, float etc)
	- empty string for string
	- false for bool
	*/

	// constants, these cannot be redeclared

	const TAX_RATE_AU float32 = 0.10
	// TAX_RATE_AU = 0.15 does not work

	fmt.Println(b,x, s4, b2, b3, s1, x1, x2, x3, f, g, h, weight, height, TAX_RATE_AU)

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

	// IF STATEMENTS
	if b {
		fmt.Println("condition met b")
	}
	if b3 {
		fmt.Println("condition met b3")
	}
	if x==0 {
		fmt.Println("condition met x")
	}

	// Shorthand if block

	if max := 10; max > 3 {
		fmt.Println("max is greater than x")
	}
	// fmt.Print(max) max is scoped to if block and cant be used here

	// SWITCH

	var myColour = "BLUE"
	switch myColour {
	case "BLUE":
		fmt.Println("found blue")
	case "RED":
		fmt.Println("found red")
	default:
		fmt.Println("default case met")
	}
}