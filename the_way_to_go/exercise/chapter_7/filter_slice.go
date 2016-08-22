package main

import (
	"fmt"
)

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
	/*output
	[0 2 4 6 8]
	*/
}
