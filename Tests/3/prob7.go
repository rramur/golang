package main

import (
	"fmt"
)

func main() {

	i := "Raghu Ram"
	if i == "Cisco" {
		fmt.Println("Found ", i)
	} else if i == "Raghu Ram" {
		fmt.Println("Found ", i)
	} else {
		fmt.Println("Noting Found ")
	}

}
