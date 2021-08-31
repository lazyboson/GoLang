package main

import (
	"errors"
	"fmt"
	"sort"
)

//specific way of naming the returning type name
//“One important thing to note: the name that’s used for a named returned value is local to the function;
//it doesn’t enforce any name outside of the function.
//It is perfectly legal to assign the return values to variables of different names:”

func divideAndRemainder(numerator, denominator int) (quotient int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Can't divide with zero")
		return 0, 0, err
	}
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder, err
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {

	//A shadowing variable is a variable that has the same name as a variable in a containing block.
	//For as long as the shadowing variable exists, you cannot access a shadowed variable.”

	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	//“You also need to be careful to ensure that you don’t shadow a package import like fmt”

	//using divsion func
	x, y, z := divideAndRemainder(4, 5)
	fmt.Println(x, y, z)

	//anonymous function -“there are two situations where declaring anonymous functions without assigning them to variables is useful: defer statements and launching goroutines”

	for i := 0; i < 10; i++ {
		func(j int) {
			fmt.Println("printing inside inline func ", j)
		}(i)
	}

	//“Functions declared inside of functions are special; they are closures”
	//“One thing that closures allow you to do is limit a function’s scope”

	type Person struct {
		firstName string
		lastName  string
		Age       int
	}

	people := []Person{
		{"ram", "shukla", 10},
		{"shyam", "pandey", 30},
		{"vivek", "kaul", 28},
	}
	fmt.Println(people)

	//sort the slice with closure function
	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	fmt.Println(people)
	//returning function from the function

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
