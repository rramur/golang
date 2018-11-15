package main

import (
	"fmt"
)

func sum(f func(n ...int) int, num ...int) int {
	x := f(num...)
	return x
}

func main() {

	c := func(n ...int) int {
		sum := 0
		for _, v := range n {
			sum += v
		}
		return sum
	}

	x := sum(c, 1, 2, 3)

	fmt.Println(x)
}
