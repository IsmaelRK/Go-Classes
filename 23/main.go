package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		prt("Go")
		waitGroup.Done() // -1
	}()

	go func() {
		prt("Function")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func prt(str string) {

	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Second)

	}

}
