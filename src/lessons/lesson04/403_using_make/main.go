package main

import "fmt"

func main() {

	/*
		A slice, being dynamic is a multistep process. Using the
		built-in function make is an efficient and optimised way
		of creating a slice as it handles the underlying array
		generation.
	*/

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// we hit the ceiling of the declared type with make
	x = append(x, 2000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// smash the doors in...
	// look at the capacity now, it has shifted from
	// 12 to 24.
	x = append(x, 3000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
