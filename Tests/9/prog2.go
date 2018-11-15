package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send_it(c, 42)

	// Not having go routene will forece the main function to wait in recv till the send is complete
	recv_it(c)

	fmt.Println("End of function")
}

// Send

func send_it(c chan<- int, v int) {

	c <- v

}

// Recieve

func recv_it(c <-chan int) {

	fmt.Println("Channel Data: \t ", <-c)

}
