package main

import "fmt"

func main() {

	fmt.Println("Enter Upper Limit:")
	var ul int
	fmt.Scanf("%d", &ul)

	fmt.Println("Enter Lower Limit:")
	var ll int
	fmt.Scanf("%d", &ll)

	fmt.Println("Result")

	for i := ll; i <= ul; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
