package main

import "fmt"

func runVariables() {

	// VARIABLES

	// declare multiple
	var f, g, h int

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

	fmt.Println(b, x, s4, b2, b3, s1, x1, x2, x3, f, g, h, weight, height, TAX_RATE_AU)



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