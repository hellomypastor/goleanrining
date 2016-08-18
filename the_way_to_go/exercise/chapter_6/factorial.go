package main

import (
	"fmt"
)

func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 1 {
		fac = n * Factorial(n-1)
		return
	}
	return
}

func main() {
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}
}
