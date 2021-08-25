package main

import (
	"fmt"
	"runtime/debug"
)

func recoverCandidateName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func CandidateName(firstName *string, lastName *string) {
	defer recoverCandidateName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from CandidateName")
}

//Recover works only when it is called from the same goroutine which is panicking.
//It's not possible to recover from a panic that has happened in a different goroutine.
// Let's understand this using an example.

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true

}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Prataysha"
	CandidateName(&firstName, nil)
	fmt.Println("returned normally from main")
	sum(5, 0)
	fmt.Println("normally returned from main")
}
