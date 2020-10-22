package main

import "fmt"

func main() {
	// loop demo with a for loop

	for i := 0; i < 10; i++ {
		if i > 0 && (i%3 == 0 || i%5 == 0) {
			continue
		}
		fmt.Printf("the value of i is: %v \n", i)

		// we have elected to loop while the counter is less than 10,
		// however another conditional counter is raised here checking
		// if the output is greater than zero and divisible by 7, if so
		// the break keyword is executed and the loop is halted earlier
		// than expected (which would have been 9)
		if i > 0 && i%7 == 0 {
			fmt.Printf("currently on %v and that is divisible by 7, so I'm breaking \n", i)
			break
		}
	}
}
