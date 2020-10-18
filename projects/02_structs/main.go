package main

import "fmt"

type contactInfo struct {
	email    string
	postcode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	emp := person{
		firstName: "Meggy",
		lastName:  "Moo",
		contactInfo: contactInfo{
			email:    "moomoo.moo.u",
			postcode: "PR116PP",
		},
	}

	emp.print()

	ptr := &emp
	ptr.updateNameValue("Ed")
	emp.print()
}

func (p *person) updateNameValue(receivedNameValue string) {
	(*p).firstName = receivedNameValue
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
