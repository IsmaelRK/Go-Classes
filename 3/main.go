package main

import "fmt"

func sum(num, num2 int8) int8 {
	r := num + num2
	return r
}

func main() {

	var (
		n1 int8 = 2
		n2 int8 = 3
	)

	val := sum(n1, n2)

	fmt.Println(val)
}
