package main

import "fmt"

var (
	z             string = "first"
	nameBoyfriend        = "harshu"
)

func main() {

	fmt.Println(z)
	z = z + "second"
	fmt.Println(z)
	z += "second"
	fmt.Println(z)
}

func two() {

	fmt.Println("My boyfriend's name is", nameBoyfriend)

}
