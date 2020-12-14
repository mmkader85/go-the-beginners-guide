package main

import "fmt"

type bot interface {
	getLanguage() string
	getGreetings() string
}

type tamBot struct{
	language string
	greeting string
}

type engBot struct{
	language string
	greeting string
}

func main() {
	tb := tamBot{
		language: "தமிழ்",
		greeting: "நல்வரவு",
	}
	eb := engBot{
		language: "English",
		greeting: "Welcome",
	}

	printGreetings(tb)
	printGreetings(eb)
}

func printGreetings(b bot) {
	fmt.Println(b.getLanguage() + ": " + b.getGreetings())
}

func (tb tamBot) getGreetings() string {
	return tb.greeting
}

func (tb tamBot) getLanguage() string {
	return tb.language
}

func (eb engBot) getGreetings() string {
	return eb.greeting
}

func (eb engBot) getLanguage() string {
	return eb.language
}
