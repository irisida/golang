package main

import "fmt"

const (
	_  = iota             // 0 is thrown away
	kb = 1 << (iota * 10) // iota1 * 10 shifts 10
	mb = 1 << (iota * 10) // iota2 * 10 shifts 20
	gb = 1 << (iota * 10) // iota3 * 10 shifts 30
)

func main() {
	fmt.Printf("%12d\t%10x\t%b \n", kb, kb, kb)
	fmt.Printf("%12d\t%10x\t%b \n", mb, mb, mb)
	fmt.Printf("%12d\t%10x\t%b \n", gb, gb, gb)
}
