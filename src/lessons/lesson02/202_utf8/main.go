package main

import "fmt"

func main() {

	s := "This is a string in UTF8"
	//bs := []byte(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U %d \n", s[i], s[i])
	}

}
