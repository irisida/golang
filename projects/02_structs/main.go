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
	/*
		pass by value demonstration. Here we will create
		a struct and show that to be able to update the
		original struct we can use pointers.
	*/

	emp := person{
		firstName: "Meggy",
		lastName:  "Moo",
		contactInfo: contactInfo{
			email:    "moomoo.moo.u",
			postcode: "PR116PP",
		},
	}

	emp.print()

	emp.updateNameValue("Ed")
	emp.print()

	/*
		reference demonstration. Here we will demo the
		mutation of a struct (slice) that does not have
		a pointer in play because a struct has an implied
		reference as part of its structure.
	*/

	mySlice := []string{"One", "Two", "Three, Four, Five"}

	fmt.Println(mySlice)
	sliceUpdater(mySlice)
	fmt.Println(mySlice)
}

func (p *person) updateNameValue(receivedNameValue string) {
	(*p).firstName = receivedNameValue
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func sliceUpdater(sl []string) {
	sl[0] = "Updated Value"
}
