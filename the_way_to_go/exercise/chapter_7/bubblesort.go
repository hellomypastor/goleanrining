package main

import (
	"fmt"
)

func bubblesort(s []int) {
	for pass := 1; pass < len(s); pass++ {
		for i := 0; i < len(s)-pass; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
}
func main() {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	bubblesort(sla)
	fmt.Println("after sort: ", sla)
}
