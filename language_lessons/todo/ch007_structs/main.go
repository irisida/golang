package main

import "fmt"

func main() {

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	var person Person

	fmt.Printf("%T \n", person)
}
