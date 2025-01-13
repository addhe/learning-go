package main

import (
	"fmt"
	"strings"
)

func main() {
	// String to be converted to uppercase.
	inputString := "Hello, World!"

	// Convert the string to uppercase using strings.ToUpper.
	uppercaseString := strings.ToUpper(inputString)

	// Print the uppercase string.
	fmt.Println(uppercaseString) 
}