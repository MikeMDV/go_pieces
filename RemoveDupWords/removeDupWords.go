package main

import (
	"fmt"
	"strings"
)

// RemDupWords will remove duplicate words in a string. The function
// will preserve the order of the words in the original string.
// The function assumes the words are separated by a single space, but the
// ability to designate a character for separation may be added later.
func RemDupWords(s string) []string {
	var uniqueWords []string

	// Trim leading and trailing whitespace from string
	s = strings.TrimSpace(s)

	// Split string on single space
	tempWords := strings.Split(s, " ")

	var wordSet = make(map[string]struct{})

	// Add only unique words to map and output slice
	for _, word := range tempWords {
		if _, ok := wordSet[word]; !ok {
			wordSet[word] = struct{}{}
			uniqueWords = append(uniqueWords, word)
		}
	}

	return uniqueWords
}

func main() {
	stringOne := "This This string string has has duplicate duplicate words words"
	fmt.Println(stringOne)
	fmt.Println(RemDupWords(stringOne))
	stringTwo := "This string also has dups This string also has dups"
	fmt.Println(stringTwo)
	fmt.Println(RemDupWords(stringTwo))

}
