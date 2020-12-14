package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	//fileName := "writer_reader.log"
	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}