package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "The quick brown fox jumps over the lazy fox. The fox is quick."
	wordToCount := "fox"

	words := strings.Fields(text)
	fmt.Printf("Words: %v\n", words)
	count := 0
	for _, word := range words {
		if strings.ToLower(word) == strings.ToLower(wordToCount) { // Case-insensitive comparison
			count++
		}
	}

	fmt.Printf("The word '%s' appears %d times in the string.\n", wordToCount, count)
}