package main

import (
	"fmt"
)

// fib will return the nth fibonacci number
func fib(n int64) int64 {
	if n == 0 {
		return 0
	}
	var a int64 = 0
	var b int64 = 1
	var i int64
	for i = 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return a + b
}

func main() {
	// This program will print the first 76 numbers in the
	// fibonacci sequence, each on a separate line. The sequence
	// starts with 0 rather than 1
	fmt.Println("The first 76 numbers in the fibonacci sequence are:")
	var i int64
	for i = 0; i < 76; i++ {
		fmt.Println(fib(i))
	}
}
