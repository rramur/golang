package main

import (
	"fmt"
)

// It's about slice and it dont rquire the size limitation

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)

}
