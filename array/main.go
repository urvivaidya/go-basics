package main

import "fmt"

func main() {
	const size = 10
	var x [size]int
	index := 4
	x[index] = 100
	fmt.Println(x)
	excercise()

}

func excercise() {
	var y [100]int
	for z := 0; z <= 99; z++ {
		y[z] = 500 + z

	}
	fmt.Println(y)
}
