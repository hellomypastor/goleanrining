package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
}

func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

func (car *Car) Start() {
	fmt.Println("Car is started")
}

func (car *Car) Stop() {
	fmt.Println("Car is stoped")
}

func (car *Car) GoToWorkIn() {
	car.Start()
	car.Stop()
}

func main() {
	m := Mercedes{Car{4}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
