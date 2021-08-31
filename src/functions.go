package main

import (
	"fmt"
	"os"
)

func average(xs []float64) float64 {
	//panic("not implemented")
	ans := 0.0
	for _, value := range xs {
		ans += value
	}
	return ans / float64(len(xs))
}

//function taking multiple int arguments implicitly known as variadic parameter -- ... is called ellipsis
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//defer, panic and recover
//defer calls the function post completion of parent function example -

func firts() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("third")
}

type Adder struct {
	start int
}

//method-- using the Adder Struct -- and can be invoked

//method vs function ??
//use methods if values need to use which are inititated at the start of the program
//functions needs to be used if ur output only depends on the supplied inputs
func (a Adder) AddTo(value int) int {
	return a.start + value
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	//passing slice with multiple arguments
	xs := []int{1, 2, 3, 4}
	fmt.Println(add(xs...))

	fmt.Println(add(1, 2, 3))

	//closure -- create function inside the function
	x := 0
	increment := func() int {
		x++
		return x
	}
	summation := func(x, y int) int {
		return x + y
	}

	fmt.Println(summation(2, 3))
	fmt.Println(increment())
	fmt.Println(increment())
	//defer is often used when source needs to be freed in some way
	defer second()
	firts()
	third()
	f, _ := os.Open("/Users/ashu/Desktop/first.txt")
	//automatically close the file handler once the job is done
	defer f.Close()

	//panic - it generally indicate the programming error or an exception condition and there is no way to recover from the panic
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
