package main

import "fmt"

func runStructs() {
	// STRUCTS
	type MapPointer struct {
		Longitude float64
		Latitude  float64
	}

	var coord MapPointer = MapPointer{144.654321, -35.321658}
	// Fields can be set explicitly or implicitly (above)
	var coord2 MapPointer = MapPointer{Longitude: 134.654321, Latitude: -32.321658}
	fmt.Println("My coordinate struct: ", coord.Latitude, coord.Longitude, coord2.Latitude, coord2.Longitude)
	var defaultCoord MapPointer = MapPointer{}
	fmt.Println("Default values for struct should be zero values (0, 0)", defaultCoord.Latitude, defaultCoord.Longitude)

	// Changing struct value via pointer
	coordPointer := &coord
	coordPointer.Latitude = 160.123
	fmt.Println(coord.Latitude)
}