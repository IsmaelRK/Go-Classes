package main

import "fmt"

func fun1() {
	fmt.Println("funcao 1")
}

func fun2() {
	defer fmt.Println("Adiado")
	fmt.Println("Aparece antes do return")
	fmt.Println("funcao 2")

	return
}

func main() {

	defer fun1()

	fun2()

}
