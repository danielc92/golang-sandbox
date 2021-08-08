package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
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

func main() {


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
	var stars = map[string]Point {
		"Alpha024": { 134134, 4135135},
		"Gamma02464": { 434134, 4135132},
		"Beta05324": { 534134, 4135133},
	}

	fmt.Println("Phonebook map: ",phoneBook, phoneBook["daniel"])
	fmt.Println("Stars map: ", stars)

	// mutating maps 

	// changing keys value
	stars["Alpha024"] = Point{ 1, 2}

	// deleting a key
	delete(stars, "Beta05324")

	// adding a key, value
	stars["New89"] = Point{888,124}
	fmt.Println("Stars map has been mutated: ", stars)
	
	// check key exists
	result, ok := stars["New89"]
	result2, ok2 := stars["New89963"]

	fmt.Println("Check stars: ", result, ok, result2, ok2)
	// STRUCTS
	type MapPointer struct {
		Longitude float64
		Latitude float64
	}

	var coord MapPointer = MapPointer{144.654321, -35.321658}
	// Fields can be set explicitly or implicitly (above)
	var coord2 MapPointer = MapPointer{Longitude: 134.654321, Latitude:-32.321658}
	fmt.Println("My coordinate struct: ", coord.Latitude, coord.Longitude, coord2.Latitude, coord2.Longitude)
	var defaultCoord MapPointer = MapPointer{}
	fmt.Println("Default values for struct should be zero values (0, 0)", defaultCoord.Latitude, defaultCoord.Longitude)

	// Changing struct value via pointer
	coordPointer := &coord
	coordPointer.Latitude = 160.123
	fmt.Println(coord.Latitude)

	// TIME MODULE
	fmt.Println("Get datetime now", time.Now())
	fmt.Println("Get unix now", time.Now().Unix()) // Note this is in seconds, not milliseconds
	fmt.Println("Get month now", time.Now().Month())
	fmt.Println("Get year now", time.Now().Year())
	fmt.Println("Adding x hours", time.Now().Add(time.Hour * 3))
	fmt.Println("Adding x minutes", time.Now().Add(time.Minute * 23))
	fmt.Println("Adding x seconds", time.Now().Add(time.Second * 50))
	fmt.Println("Get location", time.Now().Location())

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