package main

import (
	"fmt"
)

func getLocalPointer() *int {
	data := 5
	return &data
}

func modify(slice []int) {
	slice[0] = 1000
}

func main() {
	current := 100
	var pointer *int = &current //remember int* acts as declaration and assignment
	data := new(int)            // new way of creating the pointer
	data = &current             //assigning address to the pointer
	fmt.Println(*data)          //derefrencing the pointer
	*pointer++
	fmt.Println(current) //updating the value of current using the pointer

	//we can access local variable of a function in main -- it would have been error in the c++
	value := getLocalPointer()
	fmt.Println(*value)

	//always use slice instead of array pointer to change value of array inside a function for clean and bug free code
	slice := []int{
		1, 2, 3, 4, 5,
	}
	modify(slice)
	for _, value := range slice {
		fmt.Println(value)
	}

	//go doesn't support pointer Airthmatic like in c++, c
}
