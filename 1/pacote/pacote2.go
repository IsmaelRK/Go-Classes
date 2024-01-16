package pacote

import "fmt"

func write() {
	fmt.Println("Wont Write, its private")
}

func Write() {
	fmt.Println("Will Write, its public")
}
