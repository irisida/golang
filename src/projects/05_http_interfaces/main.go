package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	fmt.Println("http interfaces example")

	resp, err := http.Get("https://google.co.uk")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	/*
		behind the scenes
		data := make([]byte, 99999)
		resp.Body.Read(data)
		fmt.Println(string(data))
	*/

	/*
		io.Copy implements takes params that
		satisfy the writer interface: os.Stdout
		satisfy the reader interface: resp.body
	*/
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("just wrote %v bytes. \n", len(bs))
	return len(bs), nil
}
