package main

import (
	"fmt"
)

func main() {
	a, b := 10, 0
	c := a / b
	panic(c) //panic: runtime error: integer divide by zero
	fmt.Println(c)
}
