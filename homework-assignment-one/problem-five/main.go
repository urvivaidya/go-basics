package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)

	fmt.Println("Enter Upper Limit:")
	var ul int
	fmt.Scanf("%d", &ul)

	fmt.Println("Enter Lower Limit:")
	var ll int
	fmt.Scanf("%d", &ll)
	total := 0
	for i := 1; i <= input; i++ {
		fmt.Println("Enter number", i, ":")
		var x int
		fmt.Scanf("%d", &x)

		if x < ul && x > ll {
			total++
		}
	}
	fmt.Println("Answer:", total)
}
