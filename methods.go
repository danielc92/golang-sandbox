package main

import (
	"fmt"
	"strings"
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



type Coordinate struct {
	Longitude float64
	Latitude  float64
}

/*
	No classes exist in Go.
	However it is possible to assign a method to a type, which has a similar usage and feel to classes in other languages
	Receiver types can be used on non-struct types such as floats, ints, strings
	Usage: type.method()
	Declare: func (type T) methodName(args) T {}
*/
func (coord Coordinate) IsInBounds() bool {
	if coord.Latitude < 100 && coord.Longitude < 100 {
		return true
	}
	return false
}
func (coord Coordinate) IsInOcean() bool {
	// add some logic here
	return true
}
func (coord Coordinate) IsOnLand() bool {
	// add some logic here
	return true
}

// note: receiver must be a type, not a primitive
type MyString string
func (s *MyString) TransformUpper() {
	*s = MyString(strings.ToUpper(string(*s)))
}

func runMethods() {

	// FUNCTIONS
	fmt.Println(addThree(23,10,22))
	fmt.Println(addThree(23,10,-7))
	string1, string2 := returnMultiple()
	fmt.Println("string1", string1, "string2: ", string2)
	string3, string4 := nakedReturn()
	fmt.Println(string3, string4)


	home := Coordinate{34.23, 120.23}
	fmt.Println(home.IsInBounds(), home.IsInOcean(), home.IsOnLand())

	var customString MyString = "this should be uppercase soon~!"
	customString.TransformUpper()
	fmt.Println(customString)

	/*
	Reasons to use a pointer receiver
	- you have a large struct (avoid copying it)
	- you want to modify a variable in place
	*/

}