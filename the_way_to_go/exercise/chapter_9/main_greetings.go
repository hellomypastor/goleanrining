package main

import (
	"fmt"

	"./greetings/greetings"
)

func main() {
	name := "James"
	fmt.Println(greetings.GoodDay(name))
	fmt.Println(greetings.GoodNight(name))

	if greetings.IsAm() {
		fmt.Println("Good Morning ", name)
	} else if greetings.IsAfternoon() {
		fmt.Println("Good Afternoon", name)
	} else if greetings.IsEvening() {
		fmt.Println("Good Evening", name)
	} else {
		fmt.Println("Good Night", name)
	}
}
