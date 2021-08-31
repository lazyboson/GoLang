package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

//method declaration --“Method declarations look just like function declarations, with one addition: the receiver specification.
// The receiver appears between the keyword func and the name of the method”
//mo function or method overloading in go -- or we can't use same name
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total:%d lastUpdated:%v", c.total, c.lastUpdated)
}

func (p Person) String(test int) string {
	return fmt.Sprintf(("%s %s %d %d"), p.firstName, p.lastName, p.age, test)
}

func main() {
	p := Person{
		firstName: "ram",
		lastName:  "pandey",
		age:       21,
	}
	output := p.String(5)
	fmt.Println(output)

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
