package main

import (
	"fmt"
)

func main() {

	t := struct {
		name string
		fav  []string
		rec  map[int]string
	}{
		name: "Raghu",
		fav: []string{
			"Red",
			"Orange",
			"Yellow",
		},
		rec: map[int]string{
			1: "First Record",
			2: "Second Record",
			3: "Third Record",
		},
	}

	fmt.Println(t)

	for i, v := range t.rec {
		fmt.Println("Index: ", i, v)
	}

}
