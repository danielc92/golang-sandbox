package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func DelayedFunc() {
	fmt.Println("Daniel was here.")
}
func runAdditionalModules() {
	// TIME MODULE
	fmt.Println("Get datetime now", time.Now())
	fmt.Println("Get unix now", time.Now().Unix()) // Note this is in seconds, not milliseconds
	fmt.Println("Get month now", time.Now().Month())
	fmt.Println("Get year now", time.Now().Year())
	fmt.Println("Adding x hours", time.Now().Add(time.Hour*3))
	fmt.Println("Adding x minutes", time.Now().Add(time.Minute*23))
	fmt.Println("Adding x seconds", time.Now().Add(time.Second*50))
	fmt.Println("Get location", time.Now().Location())
	
	today := time.Now()

	if (today.Month() == time.August) {fmt.Println("correct")}

	future := time.Date(2021,12,20,10,5,1,0, time.Now().Local().Location())
	fmt.Println("Future: ",time.Until(future).Hours())

	// MATH MODULE
	fmt.Println("Finding the ceiling (rounding up)", math.Ceil(3.44))
	fmt.Println("Finding the floor (rounding down)", math.Floor(4.99))
	fmt.Println("Finding the square root", math.Sqrt(3.34))
	fmt.Println("Finding the power of 4^3", math.Pow(4, 3))
	fmt.Println("Math constants", math.Pi, math.SqrtE, math.E)
	danielsRandomNo := rand.Intn(20)
	fmt.Println("Generating random numbers", danielsRandomNo)
}