package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{5, 3}
	sq := &Square{5}
	c := &Circle{2.5}
	shapes := []Shaper{r, sq, c}
	for n, _ := range shapes {
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}
}
