package main

import (
	"fmt"
)

type vechile struct {
	colour string
	door   int
}

type truck struct {
	vechile
	four_wheel bool
}

type sedan struct {
	vechile
	luxary bool
}

func main() {

	t := truck{
		vechile: vechile{
			colour: "Yellow",
			door:   4,
		},
		four_wheel: true,
	}

	s := sedan{
		vechile: vechile{
			colour: "Red",
			door:   4,
		},
		luxary: true,
	}

	fmt.Println(t, s)
	
	fmt.Println(t.colour)
	fmt.Println(s.colour)

}