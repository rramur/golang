package main

import (
	"fmt"
)

func main() {
	a := func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}

	fmt.Println("Start of main")
	fmt.Printf("Type of a : %T\n", a)
	a()
	fmt.Println("End of main")
}