package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("https://www.google.com.sg")
	if err != nil {
		fmt.Println("Error in http.Get: " + err.Error())
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	logFile := "writer_reader.log"
	err := ioutil.WriteFile(logFile, bs, 0666)
	if err != nil {
		return 0, err
	}

	bsl := len(bs)
	fmt.Printf("%d bytes written to %s\n", bsl, logFile)

	return bsl, nil
}