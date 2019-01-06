package main

import "fmt"

func main() {
	//arrays

	//example 1

	var x [3]int

	fmt.Println(x)

	x[1] = 25

	fmt.Println(x[1])

	//example 2

	var k [3][2]float64

	fmt.Println(k)

	//example 3

	y := [5]int{1, 2, 3, 4, 5}

	fmt.Println(y)

	//example 4

	j := [...]int{1, 2, 3, 4, 5}

	fmt.Println(j)

	//example 5
	h := [...]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(h)
}
