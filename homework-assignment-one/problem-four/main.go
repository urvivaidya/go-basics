package main

import "fmt"

func main() {

	var numbers [5]int

	var total int = 0
	for i := 0; i < 5; i++ {
		fmt.Println("Enter number", i, ":")
		var x int
		fmt.Scanf("%d", &x)
		numbers[i] = x
		total += numbers[i]

	}
	fmt.Println("Answer:", total)

}
