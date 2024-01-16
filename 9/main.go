package main

import "fmt"

func main() {

	num := 5

	if num > 5 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if otherNum := num + 1; otherNum > 5 {
		fmt.Println("Yep", otherNum)
	}

	// switch

	switch num {

	case 5:
		println("Vale 5")

	case 6:
		println("Vale 6")
		return

	default:
		println("Default")

	}

}
