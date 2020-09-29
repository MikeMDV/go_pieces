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

func TestRemDupWords(t *testing.T) {
	type testCase struct {
		name         string
		inputStr     string
		expectedStrs []string
	}

	var testCases = []testCase{
		testCase{
			name:     "Test adjacent duplicated words",
			inputStr: "This This string string has has duplicate duplicate words words",
			expectedStrs: []string{
				"This",
				"string",
				"has",
				"duplicate",
				"words",
			},
		},
		testCase{
			name:     "Test duplicated words in random order",
			inputStr: "This string string words This duplicate words has duplicate has",
			expectedStrs: []string{
				"This",
				"string",
				"words",
				"duplicate",
				"has",
			},
		},
		testCase{
			name:     "Test no duplicated words",
			inputStr: "This string has no duplicate words",
			expectedStrs: []string{
				"This",
				"string",
				"has",
				"no",
				"duplicate",
				"words",
			},
		},
		testCase{
			name:     "Test string with one word",
			inputStr: "This",
			expectedStrs: []string{
				"This",
			},
		},
		testCase{
			name:     "Test empty string",
			inputStr: "",
			expectedStrs: []string{
				"",
			},
		},
	}

	for _, testCase := range testCases {
		strs := RemDupWords(testCase.inputStr)

		// Uncomment the lines below for debug purposes
		/*fmt.Println("Test name:", testCase.name)
		fmt.Println("strs received:", strs)*/

		for i, str := range strs {
			if str != testCase.expectedStrs[i] {
				t.Fatalf("Test, '%v', failed. Expected %v for a string and received %v.",
					testCase.name, testCase.expectedStrs[i], str)
			}
		}
	}
}
