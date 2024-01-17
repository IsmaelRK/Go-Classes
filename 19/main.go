package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func writeArea(f forma) {
	fmt.Printf("Area is %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type retangulo struct {
	width  float64
	height float64
}

type circle struct {
	raio float64
}

func main() {
	r := retangulo{10, 5}
	writeArea(r)

	c := circle{25}
	writeArea(c)
}
