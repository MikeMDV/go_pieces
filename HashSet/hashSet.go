package main

import (
	"fmt"
)

// A hash set can be implemented with a map to empty struct
// An empty struct in Go is 0 bytes
var hashSet = make(map[string]struct{})

func main() {
	hashSet["hello"] = struct{}{}
	hashSet["partner"] = struct{}{}
	hashSet["fancy"] = struct{}{}
	var empty struct{}
	hashSet["poodle"] = empty

	// Print the hash set
	// Expected output:
	// map[fancy:{} hello:{} partner:{} poodle:{}]
	fmt.Println(hashSet)

	// Print the key value pairs in the hash set
	// Expected output:
	// poodle {}
	// hello {}
	// partner {}
	// fancy {}
	for k, v := range hashSet {
		fmt.Println(k, v)
	}

	// Print just the keys in the hash set
	// Expected output:
	// poodle
	// hello
	// partner
	// fancy
	for k := range hashSet {
		fmt.Println(k)
	}
}
