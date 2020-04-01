package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println("*** \n\n** \n*")

	fmt.Println("32132 x 42252 =", 32132*42452)
	fmt.Println((true &&
		false) || (false && true) || !(false && false))
}
