package main

import "fmt"

func main() {

	var str string = "STRING"
	str2 := "STRING2"

	fmt.Println(str)
	fmt.Println(str2)

	var (
		x int = 1
		y int = 2
	)

	print("\n", x, "\n", y, "\n")

	myvar, myvar2 := "value", 5
	fmt.Println(myvar, myvar2)

	const myconst int = 5
	println(myconst)
}
