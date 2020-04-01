package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)

	var total int = 0
	for i := 1; i <= input; i++ {
		fmt.Println("Enter number", i, ":")
		var x int
		fmt.Scanf("%d", &x)
		total += x
	}
	fmt.Println("Answer:", total)

}
