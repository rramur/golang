package main

import (
	"fmt"
)

func get_sum() func(n ...int) int {
	x := func(n ...int) int {
		sum := 0
		for _, v := range n {
			sum += v
		}
		return sum
	}
	return x
}

func main() {

	f := get_sum()

        fmt.Printf("Type of function get_sum: %T\n",f)

	s := f(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s)
}
