package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// switch demo
		switch {
		case (i%3 == 0 && i%5 == 0):
			fmt.Printf("%v is divisible by 3 & 5 \n", i)
		case (i%3 == 0):
			fmt.Printf("%v is divisible by 3 \n", i)
		case (i%5 == 0):
			fmt.Printf("%v is divisible by 5 \n", i)
		default:
			fmt.Printf("%v is a default case\n", i)
		}

	}

}
