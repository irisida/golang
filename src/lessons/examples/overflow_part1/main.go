package main

import "fmt"

func main() {
	var (
		width  uint8 = 255
		height       = 255 // type inferred as int
	)

	width++

	fmt.Printf("width: value: %v type: %T \n", width, width)
	fmt.Printf("height: value: %v type: %T \n", height, height)

	if int(width) < height {
		fmt.Println("Height was greater")
	}

	// 1. 	Note the requirement to cast width to an int because we cannot
	// 		compare types that are not the same.
	// 2. 	Note that Go truncates an overflow and resets to its zero value
	// 		to avoid traditional buffer overflow errors and security holes.
}
