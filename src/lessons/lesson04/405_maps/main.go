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

	// is not present in the map and will therefore use the
	// associated zero values for the types.
	val, found := m["php"]
	fmt.Println(val, found)

	term := "js"

	// run a conditional test to load the
	// boolean (2nd value), then you can run the
	// block that check against a boolean value
	// running the block if true
	if v, f := m[term]; f {
		fmt.Println("conditional was found: ", v)
	}

}
