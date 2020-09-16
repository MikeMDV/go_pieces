package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return string(rs)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	rs := []rune(s)
	var isPalindrome bool = true
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}

func main() {
	// This program will print the reversed string (passed as an argument
	// when running the program) and indicate whether the string is a
	// palindrome. It will consider capital letters the same as lower
	// case letters, so "civic" will be considered a palindrome along
	// with "Civic"
	if len(os.Args) == 2 {
		fmt.Println(reverse(os.Args[1]))
		pMsg := fmt.Sprintf("'%s' is a palindrome? : %v", os.Args[1], isPalindrome(os.Args[1]))
		fmt.Println(pMsg)
	} else {
		msg := fmt.Sprintf("Usage: %s string", os.Args[0])
		fmt.Println(msg)
	}

}
