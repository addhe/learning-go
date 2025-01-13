package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "   Go is a great language.  "
	targetWord := "go"
	targetPrefix := "go"

	// 1. Convert to lowercase
	lowercaseString := strings.ToLower(inputString)

	// 2. Remove leading/trailing whitespace
	trimmedString := strings.TrimSpace(lowercaseString)

	// 3. Check if it contains the target word
	containsWord := strings.Contains(trimmedString, targetWord)

	// 4. Check if it starts with the target prefix
	startsWithPrefix := strings.HasPrefix(trimmedString, targetPrefix)

	fmt.Printf("Original string: \"%s\"\n", inputString)
	fmt.Printf("Lowercase string: \"%s\"\n", lowercaseString)
	fmt.Printf("Trimmed string: \"%s\"\n", trimmedString)
	fmt.Printf("Contains \"%s\": %t\n", targetWord, containsWord)
	fmt.Printf("Starts with \"%s\": %t\n", targetPrefix, startsWithPrefix)
}