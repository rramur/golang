package main

import (
	"fmt"
	"runtime"
)

func total(t *int, c chan int, n []int) {
	for _, v := range n {
		*t += v
	}
	// fucntion complete signal
	c <- 1
}

func main() {
	var numCPU = runtime.GOMAXPROCS(0)
	fmt.Println(numCPU)
	numCPU = 4
	n1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := [][]int{n1, n2, n3, n4}
	tt := []int{0, 0, 0, 0}
	
	c := make(chan int, numCPU)  // Buffering optional but sensible.

	for i, v := range sum {
		go total(&tt[i], c, v)
	}
	
	// Wait for all the threads to completed
	for i:= 0; i < numCPU; i++ {
		<- c
	}

	// All threads are complete, now complete the summary
	
	tot := 0
	for _, v := range tt {
		tot += v
	}

	fmt.Println("Total: ", tot)
}