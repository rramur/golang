package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send_it(c)

	for v := range c {
		fmt.Println("Channel Data: \t ", v)
	}

	fmt.Println("End of function")
}

// Send

func send_it(c chan<- int) {

	for i := 1; i <= 10; i++ {
		c <- i
	}
	// close will allow the range function on channel to proceed futher
	close(c)
}
