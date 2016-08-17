package main

import (
	"errors"
	"fmt"
	"math"
)

func Mysqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("can not do a sqrt of negative number")
	}
	return math.Sqrt(f), nil
}

func Mysqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("can not do a sqrt of negative number")
	} else {
		ret = math.Sqrt(f)
		err = nil
	}
	return ret, nil
}

func main() {
	ret1, err1 := Mysqrt(-1)
	if err1 != nil {
		fmt.Println("Error!Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok!Return values are: ", ret1, err1)
	}

	ret2, err2 := Mysqrt(5)
	if err2 != nil {
		fmt.Println("Error!Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok!Return values are: ", ret2, err2)
	}

	fmt.Println(Mysqrt2(5))
	/*output
	Error!Return values are:  NaN can not do a sqrt of negative number
	It's ok!Return values are:  2.23606797749979 <nil>
	2.23606797749979 <nil>
	*/
}
