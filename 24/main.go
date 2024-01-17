package main

import (
	"fmt"
	"time"
)

//func main() {
//
//	chanel := make(chan string)
//	go write("Hello World", chanel)
//
//	msg := <-chanel
//	fmt.Println(msg)
//
//}
//
//func write(txt string, chanel chan string) {
//
//	for i := 0; i < 5; i++ {
//
//		chanel <- txt
//		time.Sleep(time.Second)
//	}
//
//}

func main() {

	chanel := make(chan string)
	go write("Hello World", chanel)

	//for {
	//
	//	msg, opened := <-chanel
	//	fmt.Println(msg)
	//
	//	if !opened {
	//		break
	//	}
	//
	//}

	for msg := range chanel {
		fmt.Println(msg)
	}

}

func write(txt string, chanel chan string) {

	for i := 0; i < 5; i++ {

		chanel <- txt
		time.Sleep(time.Second)
	}

	close(chanel)

}
