package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello, World!"
	str2 := "hello, world!"

	// Convert both strings to lowercase for case-insensitive comparison.
	lowerStr1 := strings.ToLower(str1)
	lowerStr2 := strings.ToLower(str2)

	// Compare the lowercase strings.
	if lowerStr1 == lowerStr2 {
		fmt.Printf("The strings are equal (case-insensitive).\n")
	} else {
		fmt.Printf("The strings are not equal (case-insensitive).\n")
	}
}