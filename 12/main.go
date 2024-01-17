package main

import "fmt"

// Fibonacci recursive test

func rec1(slc []int, i1 int) {

	newNum := slc[i1] + slc[i1+1]
	slc = append(slc, newNum)

	i1++
	fmt.Println(i1)
	fmt.Println(slc)
	rec2(slc, i1)
}

func rec2(slc []int, i1 int) {

	newNum := slc[i1] + slc[i1+1]
	slc = append(slc, newNum)

	i1++
	fmt.Println(i1)

	if i1 > 25 {
		fmt.Println("Finished!")
		return
	} else {

		fmt.Println(slc)
		rec1(slc, i1)
	}

}

func main() {

	var num int = 1

	slc := []int{}
	slc = append(slc, num)
	slc = append(slc, num)

	i1 := 0

	rec1(slc, i1)
}
