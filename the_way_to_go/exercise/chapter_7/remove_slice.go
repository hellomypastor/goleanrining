package main

import (
	"fmt"
)

func RemoveStringSlice(slice []string, start, end int) []string {
	res := make([]string, len(slice)-(end-start))
	at := copy(res, slice[:start])
	copy(res[at:], slice[end:])
	return res
}

func main() {
	s := []string{"m", "n", "o", "p", "q", "r", "s"}
	res := RemoveStringSlice(s, 2, 4)
	fmt.Println(res)
}