package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer f.Close()

	buf := make([]byte, 25)
	readIndex := 0

	for {
		n, err := f.Read(buf)
		readIndex++

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			fmt.Println(readIndex, string(buf[:n]))
		}
	}
}
