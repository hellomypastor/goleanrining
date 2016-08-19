package main

import (
	"fmt"
)

/*
func Sum(arrF [4]float32) (sum float32) {
	for _, item := range arrF {
		sum += item
	}
	return
}
*/

func Sum(arrF []float32) (sum float32) {
	for _, item := range arrF {
		sum += item
	}
	return
}

func SumAndAverage(arrF []int) (int, float32) {
	sum := 0
	for _, item := range arrF {
		sum += item
	}
	return sum, float32(sum / len(arrF))
}

func main() {
	//var a = [4]float32{1.0, 2.0, 3.0, 4.0}
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of array is %f\n", Sum(a))
	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the array is: %d, and the average is: %f", sum, average)
}
