package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello, World!"
	str2 := "hello, world!"

	// Use strings.EqualFold for case-insensitive comparison.
	if strings.EqualFold(str1, str2) {
		fmt.Printf("The strings are equal (case-insensitive).\n")
	} else {
		fmt.Printf("The strings are not equal (case-insensitive).\n")
	}
}