package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orgin string = "ABC"
	//var an int
	//var newS string
	//var err error

	fmt.Printf("the size of ints is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orgin) //an is default 0
	if err != nil {
		fmt.Printf("origin %s is not an integer - exiting with error", orgin)
	}
	fmt.Printf("the integer is %d\n", an)
	an = an + 5
	newS := strconv.Itoa(an)
	fmt.Printf("the new string is %s\n", newS)
}
