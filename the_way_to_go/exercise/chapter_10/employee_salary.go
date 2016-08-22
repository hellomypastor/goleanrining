package main

import (
	"fmt"
)

type employee struct {
	salary float32
}

func (this *employee) giveRaise(pct float32) {
	this.salary += this.salary * pct
}

func main() {
	e := new(employee)
	e.salary = 10000
	e.giveRaise(0.04)
	fmt.Printf("Employee now makes %f\n", e.salary)

	e1 := &employee{10000}
	e1.giveRaise(0.04)
	fmt.Printf("Employee now makes %f\n", e1.salary)
}
