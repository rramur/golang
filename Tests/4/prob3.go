package main

import (
	"fmt"
)

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	
	// [42 43 44 45 46]
	t := a[:5]
	fmt.Println(t)

	// [47 48 49 50 51]
	t = a[5:]
	fmt.Println(t)

	// [44 45 46 47 48]
	t = a[2:7]
	fmt.Println(t)

	// [43 44 45 46 47]
	t = a[1:6]
	fmt.Println(t)

}
