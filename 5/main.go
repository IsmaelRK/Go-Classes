package main

import "fmt"

type people struct {
	nome     string
	lastName string
	age      uint8
	height   float32
}

type student struct {
	people
	course     string
	university string
}

func main() {

	fmt.Println(student{})

	p := people{"Ismael", "R", 16, 1.63}
	s := student{p, "IT", "UFSM"}

	fmt.Println(s)
	fmt.Println(s.course)
	fmt.Println(s.people.nome)

}
