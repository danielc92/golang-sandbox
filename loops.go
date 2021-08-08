package main

import (
	"fmt"
)


type Wood struct {
	Age float32
	Store string
	Density float32
	Name string
}

func runLoops() {
	// FOR LOOPS
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}


	// Iterating until condition is met
	sum := 0
	const MAX int = 1000

	for sum < MAX {
		sum += 160
		fmt.Println(sum)
	}

	// Running loops infinitely
	// for {
	// 	fmt.Println("Dont stop printing")
	// 	time.Sleep(time.Second * 1)
	// }	

	wood := []Wood{
		{ 25, "Wood & Things store", 354, "Oak"},
		{ 26, "Wood & Things store", 542, "Pine"},
		{ 32, "Wood & Things store", 764, "Cherry"},
	}

	// Using range to iterate through slice of structs
	for _, w := range wood {
		fmt.Println("Inside the loop: ", w.Age, w.Name)
	}
}