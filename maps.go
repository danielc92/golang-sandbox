package main

import "fmt"

func runMaps() {
	// MAPS

	var phoneBook map[string]string
	// maps must be initialized with make
	phoneBook = make(map[string]string)
	phoneBook["daniel"] = "04 0000 0000"
	phoneBook["jessie"] = "04 0000 0004"
	phoneBook["jimmy"] = "04 0000 0001"

	// maps with structs

	type Point struct {
		X uint
		Y uint
	}
	var stars = map[string]Point{
		"Alpha024":   {134134, 4135135},
		"Gamma02464": {434134, 4135132},
		"Beta05324":  {534134, 4135133},
	}

	fmt.Println("Phonebook map: ", phoneBook, phoneBook["daniel"])
	fmt.Println("Stars map: ", stars)

	// mutating maps

	// changing keys value
	stars["Alpha024"] = Point{1, 2}

	// deleting a key
	delete(stars, "Beta05324")

	// adding a key, value
	stars["New89"] = Point{888, 124}
	fmt.Println("Stars map has been mutated: ", stars)

	// check key exists
	result, ok := stars["New89"]
	result2, ok2 := stars["New89963"]

	fmt.Println("Check stars: ", result, ok, result2, ok2)
}