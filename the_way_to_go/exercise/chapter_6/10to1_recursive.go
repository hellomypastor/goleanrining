package main

import (
	"fmt"
)

func printRec(i int) {
	if i > 10 {
		return
	}
	printRec(i + 1)
	fmt.Printf("%d", i)
}
func main() {
	printRec(1)
}
