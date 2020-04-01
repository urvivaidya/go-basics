package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 {
			fmt.Println(x)
		}
	}
	fizzbuzz()
}

func fizzbuzz() {
	for z := 1; z <= 50; z++ {
		if z%3 == 0 && z%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if z%3 == 0 {
			fmt.Println("Fizz")
		} else if z%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(z)
		}
	}
}
