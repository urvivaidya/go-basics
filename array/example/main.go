package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += float64(x[i])
	}
	fmt.Println(total / 5)

}
