package main

import (
	"fmt"
)

func main() {

	a := map[string][]string{

		"First Key":  []string{"This is the", "FirstKey", "Value"},
		"Second Key": []string{"This is the", "SecondKey", "Value"},
	}

	fmt.Println(a)

	a["Third Key"] = []string{"This is the", "ThirdKey", "Value"}

	fmt.Println(a)

	for i, v := range a {
		fmt.Println("This is the key for ", i)

		for j, v1 := range v {
			fmt.Printf("Index: %d \t value: %v\n", j, v1)

		}

	}

	delete(a, "Third Key")

	fmt.Println(a)

	for i, v := range a {
		fmt.Println("This is the key for ", i)

		for j, v1 := range v {
			fmt.Printf("Index: %d \t value: %v\n", j, v1)

		}

	}

}
