package main

import (
	"fmt"
)

func main() {
	// 1 - use 2 nested for loops
	for i := 0; i < 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}

	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}
