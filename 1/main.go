package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"hellopac/pacote"
)

func main() {
	fmt.Println("Hi")
	pacote.Write_hello()
	pacote.Write()
	print("Hi")
	print("World")
	// print("\n\nspace")
	print("\n\n\n")

	erro := checkmail.ValidateFormat("umemail@gmail.com")
	fmt.Println(erro)

}
