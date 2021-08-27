package main

import (
	"fmt"
)

func main() {
	type person struct {
		name string
		page int
		pet  string
	}

	//assigning struct literal to a variable
	bob := person{
		"ram",
		40,
		"aprajita",
	}

	fmt.Println(bob.name)

	//anonymous structure - use var instead of type

	var nextPerson struct {
		name string
		age  int
		pet  string
	}
	nextPerson.name = "shyam"
	nextPerson.age = 35
	nextPerson.pet = "adavaita"

	//another way of declaring anonymous structure
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Printf("%s\n", pet.name)

}
