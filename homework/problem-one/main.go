package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanf("%d", &input)

	finalAnswer := 0
	for i := 1; i <= input; i++ {
		fmt.Println("Enter a number", i, ":")
		var x int
		fmt.Scanf("%d", &x)

		if x > finalAnswer {
			finalAnswer = x

		}
	}
	fmt.Println("Answer", finalAnswer)

}
