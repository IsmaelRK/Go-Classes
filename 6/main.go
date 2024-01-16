package main

import "fmt"

// pointers
func main() {

	var num int
	var point *int

	num = 5
	point = &num

	fmt.Println(point, num)
	fmt.Println(*point)
}
