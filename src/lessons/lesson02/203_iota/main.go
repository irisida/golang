package main

import "fmt"

/*
	The iota is a sequence that is automatically generated for us by
	Go for each of the constants that it is applicable to.
	Here we see a group declaration and yet the iota type is applied
	or attached to each. This is just a code efficiency in Go to avoid
	having to rewrite the same boilerplate on each line.
*/

func main() {
	const (
		SUN = iota
		MON
		TUE
		WED
		THU
		FRI
		SAT
	)

	fmt.Println(SUN, MON, TUE, WED, THU, FRI, SAT)

	// if we declare a new sequence of iotas the sequence
	// itself resets

	const (
		ZERO = iota
		ONE
		TWO
	)

	fmt.Println(ZERO, ONE, TWO)
}
