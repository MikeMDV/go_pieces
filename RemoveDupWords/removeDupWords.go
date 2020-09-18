package main

import (
	"fmt"
	"strings"
)

// remDupWords will remove duplicate words in a string. The function is not
// guaranteed to preserve the order of the words in the original string.
// The function assumes the words are separated by a single space, but the
// ability to designate a character for separation may be added later.
func remDupWords(s string) []string {
	var uniqueWords []string

	// Trim leading and trailing whitespace from string
	s = strings.TrimSpace(s)

	// Split string on single space
	tempWords := strings.Split(s, " ")

	var wordSet = make(map[string]struct{})

	// Add only unique words to map
	for _, word := range tempWords {
		if _, ok := wordSet[word]; !ok {
			wordSet[word] = struct{}{}
		}
	}

	// Add map keys to slice of strings
	for k := range wordSet {
		uniqueWords = append(uniqueWords, k)
	}

	return uniqueWords
}

func main() {
	stringOne := "This This string string has has duplicate duplicate words words"
	fmt.Println(stringOne)
	fmt.Println(remDupWords(stringOne))
	stringTwo := "This string also has dups This string also has dups"
	fmt.Println(stringTwo)
	fmt.Println(remDupWords(stringTwo))

}
