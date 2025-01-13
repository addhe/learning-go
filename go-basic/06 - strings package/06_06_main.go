package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a sample string."
	substring := "sample"

	// Find the length of the string.
	strLen := len(str)

	// Find the index of the substring.  Returns -1 if not found.
	index := strings.Index(str, substring)


	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length of string: %d\n", strLen)

	if index != -1 {
		fmt.Printf("Index of substring '%s': %d\n", substring, index)
	} else {
		fmt.Printf("Substring '%s' not found in the string.\n", substring)
	}
}