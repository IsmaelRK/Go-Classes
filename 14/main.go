package main

import "fmt"

func recoverApp() {
	fmt.Println("Recovering")

	if r := recover(); r != nil { // if there isnt anything to recover r == nil
		fmt.Println("Recovered")
	}

}
func grade(num int) {
	defer recoverApp()
	if num > 5 {
		fmt.Println("Ok!")
		return
	} else if num < 5 {
		fmt.Println("Ok")
		return
	}

	panic("NUM IS 5!")

}

func main() {

	num := 5
	grade(num)

	fmt.Println("Finished")
}
