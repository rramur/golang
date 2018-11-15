package main

import (
	"fmt"
)

type Circle struct {
	radius int
}

type Square struct {
	diag int
}

func (c Circle) area() int {
	fmt.Println("Radius of the circle ", c.radius)
	return c.radius
}

func (s Square) area() int {
	fmt.Println("Diagonal of the Square ", s.diag)
	return s.diag
}

type shape interface {
	area() int
}


func info(s shape) {
	s.area()
}

func main() {
	var a Circle
	var b Square
	
	a.radius = 10
	b.diag = 20

	info(a)
	info(b)

}
