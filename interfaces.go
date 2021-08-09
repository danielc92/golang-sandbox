package main

import (
	"fmt"
	"math"
)

/*
interface notes:
- interface DOES NOT CARE about implementation details
- struct functions that belong to an interface can have different implementations
- interface cannot use additional methods that arent included in the interface definition
- can be used to group structs in an abstract fashion
*/

type Shape interface {
	GetArea() float64
}

type Triangle struct {
	Width  float64
	Height float64
}

func (triangle Triangle) GetArea() float64 {
	return triangle.Height * triangle.Width * 0.5
}

type Square struct {
	Width float64
}

func (square Square) GetArea() float64 {
	return math.Pow(square.Width, 2)
}

func runInteraces() {

	shapes := []Shape {Square{4}, Triangle{2.3,45.4}, Triangle{20,10}}

	fmt.Println(shapes[0].GetArea())
	fmt.Println(shapes[1].GetArea())
	fmt.Println(shapes[2].GetArea())

}