package main

import (
	"fmt"
)

func main() {

	// Even, odd, quit channels

	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	go send_it(e, o, q)

	recv_it(e, o, q)

	fmt.Println("End of function")
}

// Send

func send_it(e, o, q chan<- int) {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

// Recv

// close will cause the channel to return the with value 0
// For ex: i, ok := <-q will end up returning ok as false in close case
// i as zero

func recv_it(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even Number: ", v)
		case v := <-o:
			fmt.Println("Odd  Number: ", v)
		case v := <-q:
			fmt.Println("Quit Number: ", v)
			return
		}
	}
}
