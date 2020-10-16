package main

import "fmt"

func main() {

	/* iota is basically a constant incrementor, or sequence. */
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	/* steps of increment can be determined */
	const (
		c4 = iota * 2
		c5
		c6
	)

	fmt.Println(c4, c5, c6)

	/* negative examples */
	const (
		x = -(iota)
		_
		y
		z
	)

	fmt.Println(x, y, z)

}
