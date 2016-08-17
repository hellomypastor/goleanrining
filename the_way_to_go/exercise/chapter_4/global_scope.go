package main

import (
	"fmt"
)

var a = "G"

func n() {
	fmt.Println(a)
}

func m() {
	a := "O"
	fmt.Println(a)
}

func main() {
	n()
	m()
	n()
}
