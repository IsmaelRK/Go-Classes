package main

import "fmt"

func changer_with_pointer(num *int) {

	*num = *num * -1

}

func main() {

	num := 5
	changer_with_pointer(&num)
	fmt.Println(num)

}
