package main

import (
	"fmt"
)

var a string

func f1() {
	fmt.Println(a)
}

func f2() {
	a = "O"
	fmt.Println(a)
	f1()
}

func main() {
	a = "G"
	fmt.Println(a)
	f2()
}
