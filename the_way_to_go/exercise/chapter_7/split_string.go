package main

import (
	"fmt"
)

func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func main() {
	str := "Google"
	for i := 0; i < len(str); i++ {
		a, b := Split(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}
}
