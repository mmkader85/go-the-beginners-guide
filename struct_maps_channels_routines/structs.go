package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	phone string
}

func main() {
	p := person{
		firstName:   "Mohamed",
		lastName:    "Abdul Kader",
		contactInfo: contactInfo{
			email: "mmkader85@gmail.com",
			phone: "+65 91419158",
		},
	}

	p.updateFirstName("Mohideen")
	p.print()
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}