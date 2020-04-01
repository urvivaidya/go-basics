package main

import "fmt"

func main() {
	for j := 1; j <= 10; j++ {
		switch j {
		case 0:
			fmt.Println(j, "zero")
		case 1:
			fmt.Println(j, "one")
		case 2:
			fmt.Println(j, "two")
		case 3:
			fmt.Println(j, "three")
		case 4:
			fmt.Println(j, "four")
		default:
			fmt.Println("x")
		}

	}
}
