package main

import "fmt"

func main() {

	// for i := 0; i < 11; i++ {
	// 	fmt.Printf("%d\t\t%b \n", i, i)
	// }

	b := 4
	fmt.Printf("%d\t\t%b \n", b, b)

	a := b >> 1
	fmt.Printf("%d\t\t%b \n", a, a)

	c := b << 1
	fmt.Printf("%d\t\t%b \n", c, c)

	d := b << 2
	fmt.Printf("%d\t\t%b \n", d, d)
}
