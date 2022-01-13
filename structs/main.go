package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//implicit assignment
	alex := person{"Alex", "Anderson"}

	//explicit assignment
	john := person{
		firstName: "John",
		lastName:  "Doe",
	}
	// announcing a variable with a zero value (empty string)
	var jane person
	jane.lastName = "Jane"
	jane.lastName = "Doe"

	fmt.Println(alex)
	fmt.Println(john)
	fmt.Println(jane)
}
