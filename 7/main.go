package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [5]string
	array[0] = "Pos 1"
	fmt.Println(array)

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	// Slice

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))

	slice = append(slice, 10)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
