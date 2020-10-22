package main

import "fmt"

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%12d\t%b \n", kb, kb)
	fmt.Printf("%12d\t%b \n", mb, mb)
	fmt.Printf("%12d\t%b \n", gb, gb)
}
