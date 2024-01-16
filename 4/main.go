package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func main() {

	var u user

	u.nome = "Ismael"
	u.idade = 16
	fmt.Println(u)

	usr2 := user{"Eu", 17}
	fmt.Println(usr2)

}
