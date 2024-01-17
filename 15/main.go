package main

import "fmt"

func closure() func() {
	txt := "Na closure"

	function := func() {
		fmt.Println(txt)
	}

	return function
}

func main() {
	txt := "Os main"
	fmt.Println(txt)

	newFunction := closure()
	newFunction()

}
