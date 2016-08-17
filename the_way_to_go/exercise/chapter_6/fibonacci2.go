package main

import (
	"fmt"
)

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}

func main() {
	pos := 4
	val, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, val)
	pos = 10
	val, pos = fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, val)
}
