package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	fav   []string
}

func main() {

	a := person{
		first: "Raghu",
		last:  "Ram",
		fav: []string{
			"Blue",
			"Yellow",
			"Grean",
		},
	}

	b := person{
		first: "Manu",
		last:  "D",
		fav: []string{
			"Pink",
			"Light Green",
			"Yellow",
		},
	}
	fmt.Println(a, b)

	ss := []person{a, b}

	fmt.Println(ss)

	for i, p := range ss {
		fmt.Println("Record ", i+1)
		fmt.Printf("First: %v\nLast: %v\nFav:\t\t", p.first, p.last)
		for _, v := range p.fav {
			fmt.Printf("%v\t", v)
		}
		fmt.Println(" ")

	}
}
