package main

import (
	"fmt"
)

func printInts(list ...int) {
	for _, v := range list {
		fmt.Printf("%d\n", v)
	}
}

func main() {
	printInts()
	printInts(2, 3)
	printInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
