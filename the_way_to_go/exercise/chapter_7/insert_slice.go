package main

import (
	"fmt"
)

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func main() {
	s := []string{"m", "n", "o", "p", "q", "r", "s"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0)
	fmt.Println(res)
	res = InsertStringSlice(s, in, 3)
	fmt.Println(res)
}
