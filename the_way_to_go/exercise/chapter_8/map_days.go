package main

import (
	"fmt"
)

var Days = map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednsday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}

func main() {
	fmt.Println(Days)
	flagHoliday := false
	for k, v := range Days {
		if v == "Thursday" || v == "holiday" {
			fmt.Println(v, " is the ", k, "th day in the week")
			if v == "holiday" {
				flagHoliday = true
			}
		}
	}
	if !flagHoliday {
		fmt.Println("holiday is not a day")
	}
}
