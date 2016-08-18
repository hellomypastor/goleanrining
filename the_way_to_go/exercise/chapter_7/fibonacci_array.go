package main

import (
	"fmt"
)

var fibs [50]int64

func main() {
	fibs[0] = 1
	fibs[1] = 2

	for i := 2; i < len(fibs); i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	for i := 0; i < len(fibs); i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, fibs[i])
	}
}
