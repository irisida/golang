package main

import "fmt"

func main() {
	// aliased types increase readability

	// byte = uint8
	// rune = int32 - unicode character an not a typical integer number

	var (
		byteval  byte
		uint8val uint8
	)

	// compiles without an issue becaise they are the same type
	uint8val = byteval

	fmt.Println(uint8val)

	var (
		runeval  rune
		int32val int32
	)

	// again, it works because the underlying type is the same
	runeval = int32val

	fmt.Println(runeval)
}
