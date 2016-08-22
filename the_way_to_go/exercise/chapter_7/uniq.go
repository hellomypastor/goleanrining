package main

import (
	"fmt"
)

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'd', 'e', 'f'}

func main() {
	arru := make([]byte, len(arr))
	ixu := 0
	temp := byte(0)

	for _, val := range arr {
		if val != temp {
			arru[ixu] = val
			ixu++
		}
		temp = val
	}
	fmt.Println(string(arru))
}
