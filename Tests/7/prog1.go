package main

import (
	"fmt"
)

type Person struct {
	name  string
	alive bool
}

func change_name(p *Person, n string) {
	(*p).name = n
}

func main() {
	a := Person{
		alive: true,
	}

	fmt.Println(a)

	change_name(&a, "Raghu Ram")

	fmt.Println(a)

}
