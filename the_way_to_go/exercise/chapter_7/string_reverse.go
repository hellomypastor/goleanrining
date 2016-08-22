package main

import (
	"fmt"
)

func reverse1(s string) string {
	s1 := []byte(s)
	var rev [100]byte
	j := 0
	for i := len(s1) - 1; i >= 0; i-- {
		rev[j] = s1[i]
		j++
	}
	return string(rev[:])
}

func reverse2(s string) string {
	s1 := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	return string(s1)
}

func reverse3(s string) string {
	runes := []rune(s)
	n, h := len(s), len(s)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	str := "Google"
	str1 := reverse1(str)
	fmt.Println(str, "--->", str1)
	str2 := reverse2(str)
	fmt.Println(str, "--->", str2)
	str3 := reverse3(str)
	fmt.Println(str, "--->", str3)
}
