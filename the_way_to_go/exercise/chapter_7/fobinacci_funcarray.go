package main

import (
	"fmt"
)

var term = 15

func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0] = 1
	farr[1] = 1
	for i := 2; i < len(farr); i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}

func main() {
	result := fibarray(term)
	for ix, fib := range result {
		fmt.Printf("The %d-th Fibonacii number is: %d\n", ix, fib)
	}
}
