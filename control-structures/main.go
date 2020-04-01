package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	two()
	three()

}
func two() {
	for h := 1; h <= 5; h++ {
		if h%2 == 0 {
			fmt.Println(h, "even")
		} else {
			fmt.Println(h, "odd")
		}
	}
}

func three() {
	for x := 25; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Println(x)
			// divisible by 2
		} else if x%5 == 0 {
			fmt.Println(x)
			// divisible by 5
		} else if x%7 == 0 {
			fmt.Println(x)
		} else {
			fmt.Println("<3")
		}

	}
}
