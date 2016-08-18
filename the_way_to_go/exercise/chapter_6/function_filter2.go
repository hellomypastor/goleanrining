package main

import (
	"fmt"
)

type flt func(int) bool

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(s1 []int, f flt) (yes, no []int) {
	for _, val := range s1 {
		if f(val) {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	even, odd := filter(slice, isEven)
	fmt.Println("the even elements of slice are: ", even)
	fmt.Println("the odd elements of sliece are: ", odd)
}
