package main

import "fmt"

func main() {

	/*
		maps demo
		maps are key:value pairs
		suitable for lookup storage and speed optimised for
		retrieval.
	*/

	m := map[string]int{
		"go":     1000,
		"js":     32500,
		"ts":     3099,
		"svelte": 5977,
	}

	fmt.Println(m)       // prints a map object
	fmt.Println(m["go"]) // prints the value associated with the key

	val, found := m["php"]
	fmt.Println(val, found)

}
