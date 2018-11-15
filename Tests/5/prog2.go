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

	m := map[string][]string{
		a.last: a.fav,
		b.last: b.fav,
	}

	for i, mv := range m {
		fmt.Println("Last: ", i)
		for _, vv := range mv {
			fmt.Printf("%v\t", vv)
		}
		fmt.Println(" ")

	}

	mm := map[string]person{
		a.last: a,
		b.last: b,
	}

	for ii, mmv := range mm {
		fmt.Println("Last: ", ii)
		fmt.Println("Person Record : ", mmv)

	}
}
