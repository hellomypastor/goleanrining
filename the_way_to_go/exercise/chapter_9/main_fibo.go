package main

import (
	"fmt"
	"test/fibo"
)

var op string
var nextFibo int

func next() {
	result := 0
	nextFibo++
	result = fibo.Fibonacci(op, nextFibo)
	fmt.Printf("fibonacci(%d) is: %d\n", nextFibo, result)
}

func calls() {
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
}

func main() {
	op = "+"
	calls()
	fmt.Println("Change of operation from + to *")
	nextFibo = 0
	op = "*"
	calls()
}
