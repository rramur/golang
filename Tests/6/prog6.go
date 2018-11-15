package main

import (
	"fmt"
)

func test() func() int {
	n := 0
	return func() int {
		n++
		return n
	}

}

func main() {

	a := test()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	b := test()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}
