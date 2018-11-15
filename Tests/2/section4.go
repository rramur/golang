package main

import (
	"fmt"
)

const (
	_  = iota
	y1 = iota + 2014
	y2 = iota + 2014
	y3 = iota + 2014
	y4 = iota + 2014
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}