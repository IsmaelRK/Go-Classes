package main

import "fmt"

func main() {

	chanel := make(chan string, 2) // buffer

	chanel <- "Hello World"
	chanel <- "Hello World2"
	msg := <-chanel

	fmt.Println(msg)

}
