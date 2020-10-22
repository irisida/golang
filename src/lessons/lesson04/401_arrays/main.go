package main

import "fmt"

func main() {
	/*
		Arrays,
		arrays are of fixed size in Go. If you're looking for a dynamic
		sized version of an array like structure you need to look at
		slices. Arrays can store a single type therefore an array must
		know both the fixed size and type at compile time.

		The array is a zero indexed structure, meaning the position zero
		is counted and the first element will be located at arrayname[0]

		most idiomatic Go code will use slices instead. In fact the lang
		docs actually say you should use slices instead but nonetheless
		it helps to see and know how to use an array structure.
	*/

	var x [10]int

	fmt.Printf("%#v \n", x)

	x[9] = 100
	fmt.Printf("%#v \n", x)
}
