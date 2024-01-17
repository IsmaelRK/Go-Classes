package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u *user) save() { // global, using *
	u.name = u.name + "-lastName"

	fmt.Println("Saving user in Db", u.name)
}

//func (u user) save() {  // Local, not using *
//	u.name = u.name + "-lastName"
//
//	fmt.Println("Saving user in Db", u.name)
//}

func main() {

	usr1 := user{"Ismael", 26}
	fmt.Println(usr1)

	usr1.save()

	fmt.Println(usr1)
}
