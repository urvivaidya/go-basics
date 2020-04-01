package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fmt.Println("Enter number", i, ":")
		var x int
		fmt.Scanf("%d", &x)
		if x%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}
