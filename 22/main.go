package main

import (
	"fmt"
	"time"
)

func main() {

	go prt("GoRoutine")
	prt("Regular Function")

}

func prt(str string) {

	for {
		fmt.Println(str)
		time.Sleep(time.Second)
	}

}
