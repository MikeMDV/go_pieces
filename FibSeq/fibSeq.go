package main

import (
	"fmt"
)

// Fib will return the nth fibonacci number
// The function will run in O(n) time since it uses
// a dynamic programming bottom up approach and memoization
// of the previous 2 numbers in the series
func Fib(n int64) (int64, error) {
	var err error
	if n < 0 {
		err = fmt.Errorf("Error: n is less than 0")
		return 0, err
	}
	if n == 0 {
		return 0, err
	}
	var a int64 = 0
	var b int64 = 1
	var i int64
	for i = 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return a + b, err
}

func main() {
	// This program will print the first 76 numbers in the
	// fibonacci sequence, each on a separate line. The sequence
	// starts with 0 rather than 1
	fmt.Println("The first 76 numbers in the fibonacci sequence are:")
	var i int64
	for i = 0; i < 76; i++ {
		fib, _ := Fib(i)
		fmt.Println(fib)
	}
}
