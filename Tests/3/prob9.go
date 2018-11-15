package main

import (
	"fmt"
)

func main() {
	s := "favSport"
	switch s {
	case "some thing":
		fmt.Println("some thing")
	case "not even close":
		fmt.Println("not even close")
	case "favSport":
		fmt.Println("This is my fav Sport")
	default:
		fmt.Println("None of the above")

	}

}
