package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "The quick brown dog jumps over the lazy fox."
	oldWord := "dog"
	newWord := "fox"

	// Replace all occurrences of oldWord with newWord.
	replacedString := strings.ReplaceAll(inputString, oldWord, newWord)

	fmt.Printf("Original string: %s\n", inputString)
	fmt.Printf("Replaced string: %s\n", replacedString)
}