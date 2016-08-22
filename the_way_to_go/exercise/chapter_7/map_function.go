package main

import (
	"fmt"
)

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}
	fmt.Println(mapFunc(mf, list))
}
