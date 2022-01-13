package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // We may just use a type alone omitting a property.
}

func main() {
	//implicit assignment
	alex := person{"Alex", "Anderson", contactInfo{"alex@exampl.ecom", 12345}}

	//explicit assignment
	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "john@example.com",
			zipCode: 12345,
		},
	}
	// announcing a variable with a zero value (empty string)
	var jane person
	jane.lastName = "Jane"
	jane.lastName = "Doe"
	jane.contact.email = "janedoe@example.com"
	jane.contact.zipCode = 12345

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@example.com",
			zipCode: 12345,
		},
	}

	alex.print()
	john.print()
	jane.print()
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
