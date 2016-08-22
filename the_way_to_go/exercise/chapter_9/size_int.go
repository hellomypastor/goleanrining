package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Printf("the size of an int is %d", size)
}
