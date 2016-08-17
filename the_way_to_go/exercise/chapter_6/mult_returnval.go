package main

import (
	"fmt"
)

func sumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func sumProductDiffN(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return s, p, d
}

func main() {
	sum, prod, diff := sumProductDiff(3, 4)
	fmt.Println("Sum: ", sum, " | Product: ", prod, " | Diff: ", diff)
	sum, prod, diff = sumProductDiffN(3, 4)
	fmt.Println("Sum: ", sum, " | Product: ", prod, " | Diff: ", diff)
	/*output
	Sum:  7  | Product:  12  | Diff:  -1
	Sum:  7  | Product:  12  | Diff:  -1
	*/
}
