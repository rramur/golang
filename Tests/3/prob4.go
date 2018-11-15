package main

import (
	"fmt"
)

func main() {

	i := 1978
	for {

		if i > 2018 {
			break
		}
		fmt.Println("I am alive at ", i)
		i++

	}

}
