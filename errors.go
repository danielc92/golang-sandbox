package main

import (
	"errors"
	"fmt"
	"log"
)

func ValidateOrError(s string) (error) {
	if len(s) > 10 {
		return errors.New("string cannot be over 10 chars")
	}

	return nil
}

func runErrors() {
	name := "Daniel W. Corcoran"
	err := ValidateOrError(name)

	if (err != nil) {
		log.Fatal(fmt.Println(err))
	}
}