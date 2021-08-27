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

	//“If you have a sparse array (an array where most elements are set to their zero value),
	// you can specify only the indices with values in the array literal:
	var ArrayData = [12]int{1, 5: 4, 6, 10: 100, 1}
	for _, _v := range ArrayData {
		fmt.Println(_v)
	}

	//another way of declaring the array
	var newArray = [...]int{1, 2, 3, 4}
	for _, _v := range newArray {
		fmt.Println(_v)
	}

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
	//“Using […] makes an array. Using [] makes a slice.”

	var slice1 []float64
	fmt.Println(len(slice1))
	//use make function to create slice
	xx := make([]float64, 5)
	fmt.Println(len(xx))

	xx = append(xx, 40)
	//fmt print len again to check if size has grown by one
	fmt.Println(xx)

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

	//create a slice using make function with zero length and non-zero capacity
	x_vec := make([]int64, 0, 100)

	// this method helps in unnecessary coping of elements as append increases the capacity twice post index increases to five

	x_vec = append(x_vec, 100, 300, 400, 500)
	x_vec = append(x_vec, 200)

	for _, line := range x_vec {
		fmt.Println(line)
	}

	//“When you take a slice from a slice, you are not making a copy of the data.
	//Instead, you now have two variables that are sharing memory.”

	yy := x_vec[1:3]
	for _, line := range yy {
		fmt.Println(line)
	}

	//modification of one will modify other

	yy[0] = 20000

	for _, line := range x_vec {
		fmt.Println(line)
	}

	//we can use inbuilt copy function for creating independent slice
	//“The copy function takes two parameters. The first is the destination slice and the second is the source slice”

	y_x := []int{1, 2, 3, 4}
	yy_x := make([]int, 4)
	num := copy(yy_x, y_x)
	fmt.Println(y, num)

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
