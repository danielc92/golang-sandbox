package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// exercise https://tour.golang.org/moretypes/23
func WordCount(s string) map[string]int {
	
	// declare and initialize the map
	m := make(map[string]int)
	
	// split string using space as a delimiter
	split := strings.Split(s, " ")
	
	// iterate through individual terms
	for _, value := range(split) {
		_, hasKey := m[value]
		// if key exist populate with 1
		if hasKey {
			m[value] = m[value] + 1
		} else {
			// if not increment
			m[value] = 1
		}
	}
	
	// return the map
	return m

	/*
	PASS
	f("I am learning Go!") = 
	map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
	PASS
	f("The quick brown fox jumped over the lazy dog.") = 
	map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
	PASS
	f("I ate a donut. Then I ate another donut.") = 
	map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
	PASS
	f("A man a plan a canal panama.") = 
	map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
  */
}

func runSlices() {
	// SLICES
	// Similar to arrays, yet dynamically sized
	var animals []string
	animals = append(animals, "Donkey")
	animals = append(animals, "Frog")
	animals = append(animals, "Leopard")
	animals = append(animals, "Rhino")
	fmt.Println("Animals: ", animals, len(animals), animals[2], animals[1:2])


	// shorthand new slice
	shortSlice := []string{"carrot cake", "shortbread"}
	fmt.Println(shortSlice)

	// slice of structs
	structSlice := []struct {
		Age int 
		Name string
	}{
		{24, "Sammy"},
		{25, "Rachel"},
		{78, "Jessica"},
	}

	fmt.Println("struct slice: ", structSlice)

	// nil slices
	var emptySlice []string
	if emptySlice == nil {
		fmt.Println("Nil slice.")
	} else {
		fmt.Println("Not empty.")
	}
	emptySlice = append(emptySlice, "item")
	if emptySlice == nil {
		fmt.Println("Nil slice.")
	} else {
		fmt.Println("Not empty.")
	}

	// using make to create slices

	slice := make([]int, 10)
	fmt.Println(slice, "slice made with 'make'")

	// image exercise
	canvas := [][]uint8{
		[]uint8{0,0,0,0},
		[]uint8{0,0,0,0},
		[]uint8{0,0,0,0},
	}
	canvasWidth := 3
	canvasHeight := 4
	for w := 0; w < canvasWidth; w++ {
		for h := 0; h < canvasHeight; h++ {
			canvas[w][h] = uint8(rand.Intn(250))
		}
	}
	fmt.Println(canvas, canvasWidth, canvasHeight)
}