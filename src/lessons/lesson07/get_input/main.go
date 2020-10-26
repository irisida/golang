package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)
	fmt.Printf("%#v \n", os.Args)

	for i := 0; i < argCount; i++ {
		fmt.Println(os.Args[i])
	}

}
