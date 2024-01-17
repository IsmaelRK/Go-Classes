package main

import "fmt"

func sum(nums ...int) (tot int) {
	tot = 0

	fmt.Println(nums)
	for _, num := range nums {
		tot += num
	}

	return
}

func main() {

	res := sum(1, 2, 3, 4, 6, 6, 78, 9, 100)
	fmt.Println(res)

	func(txt string) {
		fmt.Println("Anonima", txt)
	}("Param")

	returned := func(txt string) string {
		return fmt.Sprintf("Received --> %s", txt)
	}("Start")

	fmt.Println(returned)
}
