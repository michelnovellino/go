package main

import "fmt"

//formas de ejecutar un ciclo for
func main() {

	var i int
	i = 0

	for i <= 9 {
		fmt.Println(i + 1)
		i++
	}

	for j := 0; j <= 9; i++ {
		fmt.Println(j + 1)
	}

}
