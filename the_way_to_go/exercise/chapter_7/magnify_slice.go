// test project test.go
package main

import (
	"fmt"
)

var s []int

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	s = ns
	return s
}

func main() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
	/*output
	The length of s before enlarging is: 3
	[1 2 3]
	The length of s after enlarging is: 15
	[1 2 3 0 0 0 0 0 0 0 0 0 0 0 0]
	*/
}
