package main

import (
	"fmt"
)

func main() {
	var x [5]float64 //intialization of an array is reverse of c with type along with inclusion of var at front
	x[0] = 10
	x[1] = 10
	x[2] = 10
	x[3] = 10
	x[4] = 10
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	fmt.Println(sum / float64(len(x)))

	//go compiler don't allow us to create a variable that is non used
	//for able to use range function of the go, use following code
	var total float64 = 0.0
	for _, value := range x {
		fmt.Println(value)
		total += (float64(value))
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(x)))

	//list way of initializing array
	data := [4]int{1, 2, 34, 90}
	for _, _data := range data {
		fmt.Println(_data)
	}

	//slice has variable length - size parameter is not explicitly required in case of the slice vis-a-vis array
	var slice1 []float64
	fmt.Println(len(slice1))
	//use make function to create slice
	xx := make([]float64, 5)
	fmt.Println(len(xx))

	xxx := make([]float64, 5, 10)
	for _, _x := range xxx {
		fmt.Println(_x)
	}

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:5]
	y1 := arr[0:2]
	for _, _y := range y {
		fmt.Println(_y)
	}
	for _, _y1 := range y1 {
		fmt.Println(_y1)
	}

	//append adds the elements at the end of slice
	slice2 := []int{1, 2, 3}
	slice3 := append(slice2, 4, 5)
	fmt.Println(slice1, slice2, slice3)
	//map - key, value pair - map has to initalized before prioir to its use
	_map := make(map[string]int)
	_map["key"] = 10
	fmt.Println(_map["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berelium"
	elements["B"] = "Boron"

	fmt.Println(elements["B"])

	//looking up an element that may not exists
	name, ok := elements["in"] //first variable holds the value returned and second variable holds whether value is present or not in bool terms
	fmt.Println(name, ok)
	if name, ok := elements["um"]; ok {
		fmt.Println(name, ok)
	}

	//shoter way of creating the map
	s_map := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Pb": "Lead",
		"Ne": "Neon",
	}
	//iterating over the entire map
	for k, v := range s_map {
		fmt.Println(k, v)
	}
	//program to find smallest of the given list
	list := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	max_number := 0
	for _, value := range list {
		if value > max_number {
			max_number = value
		}
	}
	fmt.Println(max_number)
	smallest_number := list[0]
	for _, value := range list {
		if value < smallest_number {
			smallest_number = value
		}
	}
	fmt.Println(smallest_number)
}
