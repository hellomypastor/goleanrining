package main

import (
	"fmt"
	"math"
)

func maxSlice(s1 []int) (max int) {
	for _, v := range s1 {
		if v > max {
			max = v
		}
	}
	return
}

func minSlice(s1 []int) (min int) {
	min = math.MaxInt32
	for _, v := range s1 {
		if v < min {
			min = v
		}
	}
	return
}

func main() {
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	fmt.Printf("The maximum is %d\n", max)
	min := minSlice(sl1)
	fmt.Printf("The minimum is %d\n", min)
	/*output
	The maximum is 643
	The minimum is 2
	*/
}
