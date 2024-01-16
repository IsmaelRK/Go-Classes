package main

import "fmt"

func main() {

	//i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("plus", i)
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)
	//
	//for j := 0; j < 10; j++ {
	//	fmt.Println("plus", j)
	//}

	names := [3]string{"Ismael", "Joao", "Pedro"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range "WORD" {
		fmt.Println(name, string(name))
	}


	user := 
}
