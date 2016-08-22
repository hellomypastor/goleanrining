package main

import (
	"fmt"
)

func Split2(s string) string {
	mid := len(s) / 2
	return s[0:mid] + s[mid:]
}

func main() {
	str := "Google"
	str2 := Split2(str)
	fmt.Println("The string %s tranformed is %s\n", str, str2)
}
