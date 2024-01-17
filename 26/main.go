package main

import (
	"fmt"
	"time"
)

func main() {

	chanel, chanel2 := make(chan string), make(chan string)

	go func() {

		for {
			time.Sleep(time.Millisecond * 500)
			chanel <- "C1"
		}

	}()

	go func() {

		for {
			time.Sleep(time.Second * 2)
			chanel2 <- "C2"
		}

	}()

	//for {
	//	msg1 := <-chanel
	//	msg2 := <-chanel2
	//
	//	fmt.Println(msg1)
	//	fmt.Println(msg2)
	//
	//}

	for {
		select {

		case msg := <-chanel:
			fmt.Println(msg)

		case msg2 := <-chanel2:
			fmt.Println(msg2)
		}
	}

}
