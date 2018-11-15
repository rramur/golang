package main

import (
	"fmt"
)

func main() {
	a := [][]string{
		{"This", "is", "the", "Header"},
		{"This", "is", "the", "Record"},
	}

	fmt.Println(a)

	s1 := []string{"This", "is", "the", "Header"}
	s2 := []string{"This", "is", "the", "Record"}

	fmt.Println(s1)
	fmt.Println(s2)

	ss := [][]string{s1, s2}

	fmt.Println(ss)

	for i, main_s := range ss {
		fmt.Println("Header ", i)
		fmt.Println(main_s)
		for j, sub_s := range main_s {
			fmt.Println("Record ", j)
			fmt.Println(sub_s)
		}
	}

}