package main

import (
	"errors"
	"testing"
)

func TestFib(t *testing.T) {
	type testCase struct {
		name        string
		n           int64
		expectedFib int64
		expectedErr error
	}

	var testCases = []testCase{
		testCase{
			name:        "Test 0th fib number",
			n:           0,
			expectedFib: 0,
			expectedErr: nil,
		},
		testCase{
			name:        "Test 1st fib number",
			n:           1,
			expectedFib: 1,
			expectedErr: nil,
		},
		testCase{
			name:        "Test 2nd fib number",
			n:           2,
			expectedFib: 1,
			expectedErr: nil,
		},
		testCase{
			name:        "Test 3rd fib number",
			n:           3,
			expectedFib: 2,
			expectedErr: nil,
		},
		testCase{
			name:        "Test 4th fib number",
			n:           4,
			expectedFib: 3,
			expectedErr: nil,
		},
		testCase{
			name:        "Test 75th fib number",
			n:           4,
			expectedFib: 2111485077978050,
			expectedErr: nil,
		},
		testCase{
			name:        "Test negative number for n",
			n:           -1,
			expectedFib: 0,
			expectedErr: errors.New("Error: n is less than 0"),
		},
	}

	for _, testCase := range testCases {
		fib, err := Fib(testCase.n)

		// Uncomment the lines below for debug purposes
		/*fmt.Println("Test name:", testCase.name)
		fmt.Println("fib received:", fib)
		fmt.Println("err received:", err)*/

		if fib != testCase.expectedFib && err != testCase.expectedErr {
			t.Fatalf("Test, '%v', failed. Expected %d for fib number and received %d. Expected %v for err and received %v.",
				testCase.name, testCase.expectedFib, fib, testCase.expectedErr, err)
		}
	}
}
