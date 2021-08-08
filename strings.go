package main

import (
	"fmt"
	"strings"
)


func runStrings() {
	fmt.Println("Strings")

	dirtyString := "  DaNiel CoRCORA;N   "
	cleanString1 := strings.TrimSpace(dirtyString)
	cleanString2 := strings.ReplaceAll(cleanString1, ";", "")
	cleanString3 := strings.ToTitle(cleanString2)

	fmt.Println("Cleaned strings: ", cleanString1, cleanString2, cleanString3, strings.Fields("my old cat is old"))

	tokens := []string{"my orange fat cat"}
	joined := strings.Join(tokens, "_")
	fmt.Println("Tokens joined: ", joined)
}