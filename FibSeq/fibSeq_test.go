package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	// Empty for these tests
}

func tearDown() {
	// Empty for these tests
}

func TestFib(t *testing.T) {
	type testCase struct {
		name        string
		n           int64
		expectedFib int64
	}

	var testCases = []testCase{
		testCase{
			name:        "Test 0th fib number",
			n:           0,
			expectedFib: 0,
		},
		testCase{
			name:        "Test 1st fib number",
			n:           1,
			expectedFib: 1,
		},
		testCase{
			name:        "Test 2nd fib number",
			n:           2,
			expectedFib: 1,
		},
		testCase{
			name:        "Test 3rd fib number",
			n:           3,
			expectedFib: 2,
		},
		testCase{
			name:        "Test 4th fib number",
			n:           4,
			expectedFib: 3,
		},
		testCase{
			name:        "Test 75th fib number",
			n:           75,
			expectedFib: 2111485077978050,
		},
		testCase{
			name:        "Test negative number for n",
			n:           -1,
			expectedFib: 0,
		},
	}

	for i, testCase := range testCases {
		fib, err := Fib(testCase.n)

		// Uncomment the lines below for debug purposes
		/*fmt.Println("Test name:", testCase.name)
		fmt.Println("fib received:", fib)
		fmt.Println("err received:", err)*/

		if i < len(testCases)-1 {
			if fib != testCase.expectedFib || err != nil {
				t.Fatalf("Test, '%v', failed. Expected %d for fib number and received %d. Received err, %v.",
					testCase.name, testCase.expectedFib, fib, err)
			}
		} else if fib != testCase.expectedFib || err == nil {
			t.Fatalf("Test, '%v', failed. Expected %d for fib number and received %d. Received err, %v.",
				testCase.name, testCase.expectedFib, fib, err)
		}
	}
}
