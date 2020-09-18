package main

import (
	"fmt"
	"strings"
)

// remDupWords will reverse the order of words in a string. The function
// assumes the words are separated by a single space, but the ability
// to designate a character for separation may be added later.
func revWords(s string) string {
	// Trim leading and trailing whitespace from string
	s = strings.TrimSpace(s)

	// Split string on single space
	tempWords := strings.Split(s, " ")

	// Add words in reverse order to string
	// Using strings.Builder minimizes memory copying
	var b strings.Builder
	for i := len(tempWords) - 1; i >= 0; i-- {
		b.WriteString(tempWords[i])
		if i > 0 {
			b.WriteString(" ")
		}
	}

	return b.String()
}

func main() {
	stringOne := "This is a string of words"
	fmt.Println(stringOne)
	fmt.Println(revWords(stringOne))
}
