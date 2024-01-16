package main

type people struct {
	nome     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	people
	course     string
	university string
}
