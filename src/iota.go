package main

import (
	"fmt"
)

type BitField int

//“The first constant in the const block has the type specified and its value is set to iota. E
//very subsequent line has neither the type nor a value assigned to it. When the Go compiler sees this,
//it repeats the type and the assignment to all of the subsequent constants in the block”
type MailCategory int

const (
	//starts with zero
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

//iota plays similar role as enum
const (
	Field1 BitField = iota
	Field2
	Field3
	Field4
)

func main() {
	fmt.Println(Field4)
	fmt.Println(Personal)
}
