package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func runEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("something went wrong! failed to load environment from .env")
	}

	val := os.Getenv("SPIDERMAN")
	val2 := os.Getenv("HULK")

	fmt.Println(val, val2)
}