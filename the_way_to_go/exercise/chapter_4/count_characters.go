package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//count number of characters
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("the number of bytes in str1 is %d\n", len(str1))
	fmt.Printf("the number of characters in str1 is %d\n", utf8.RuneCountInString(str1))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("the number of bytes in str2 is %d\n", len(str2))
	fmt.Printf("the number of bytes in str2 is %d\n", len(str2))

	/*output
	the number of bytes in str1 is 22
	the number of characters in str1 is 22
	the number of bytes in str2 is 28
	the number of bytes in str2 is 28
	*/

}
