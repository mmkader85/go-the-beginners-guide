package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Enter path to filename to be read in the command line while running.
// $ go run main.go writer_reader.log
func main() {
	args := os.Args
	// fileName := "writer_reader.log"
	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
	}

	written, err := io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatalln("Unable to read out the file", err)
	}

	fmt.Printf("Written %d characters", written)
}