package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


func runJwt() {
	// Signing tokens
	hmacSecret := []byte("MySecret123")

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": "daniel83_4",
		"exp": time.Now().UTC().Add(time.Hour * 8).Unix(),
		"iat": time.Now().UTC().Unix(),
		})

	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println("Created: ",tokenString)
	if (err != nil) {

		fmt.Println("Token: ", tokenString)
	} else {
		fmt.Println(err)
	}

	// Validating tokens

	tokenToValidate := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg1MzY4NDYsImlhdCI6MTYyODUwODA0NiwidXNlcm5hbWUiOiJkYW5pZWw4M180In0.47POPHDdG1uMLiUmAapgGTXvSKPSEQP4uARAPldOVBBIQ4CbjUVOsdoFqvCg3IMEc6VAn5LjFSGW-CbuiA1L3g"

	value, err2 := jwt.Parse(tokenToValidate, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, err3 := token.Method.(*jwt.SigningMethodHMAC); !err3 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if (err2 == nil) {
		fmt.Println("Success! Validation results: ", value.Valid)
	}



}