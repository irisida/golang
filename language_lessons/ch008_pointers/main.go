package main

import "fmt"

func main() {

	var x int = 2

	ptr := &x

	fmt.Printf("ptr is of type %T with value of %v\n", ptr, ptr)
}
