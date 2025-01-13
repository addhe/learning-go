package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "This is a sample sentence."
	delimiter := "-"

	// Split the sentence into words using spaces as the delimiter.
	words := strings.Fields(sentence)

	// Join the words back together using the new delimiter.
	joinedSentence := strings.Join(words, delimiter)

	fmt.Printf("Original sentence: %s\n", sentence)
	fmt.Printf("Joined sentence: %s\n", joinedSentence)
}