package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Alive bool
}

func main() {
	a := person{
		First: "Raghu",
		Last:  "Ram",
		Alive: true,
	}

	b := person{
		First: "A",
		Last:  "H",
		Alive: false,
	}

	p := []person{a, b}

	fmt.Println(a, b)

	bs, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Json Raw Format\n", bs)
		fmt.Println("Json String Format\n", string(bs))
	}
}
