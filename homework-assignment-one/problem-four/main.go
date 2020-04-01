package main

import "fmt"

func main() {
	fmt.Println("Enter number of numbers:")
	var input int
	fmt.Scanf("%d", &input)
	var numbers = make([]int, input)

	var total int = 0
	for i := 0; i < input; i++ {
		fmt.Println("Enter number", i, ":")
		var x int
		fmt.Scanf("%d", &x)
		numbers[i] = x
		total += numbers[i]

	}
	fmt.Println("Answer:", total)

}
