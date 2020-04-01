package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fmt.Println("Enter a number", i, ":")
		var x int
		fmt.Scanf("%d", &x)
	}

}
