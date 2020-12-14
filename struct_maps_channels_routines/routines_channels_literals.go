package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://youtube.com",
		"https://stackoverflow.com",
		"https://linkedin.com",
		"https://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkWebLink(link, c)
	}

	//Infinite loop
	//for {
	//	// This will be called immediately once a value is received in channel
	//	go checkWebLink(<-c, c)
	//}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkWebLink(link, c)
		}(l)
	}
}

func checkWebLink(l string, c chan string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l + " is down!")
		c <- l
		return
	}

	fmt.Println(l + " is up!")
	c <- l
}
