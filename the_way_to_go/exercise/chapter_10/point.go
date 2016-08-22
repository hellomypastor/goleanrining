package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Pola struct {
	R, T float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func Scale(p *Point, s float64) (q Point) {
	q.X = p.X * s
	q.Y = p.Y * s
	return
}

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector is %f\n", Abs(p1))

	p2 := &Point{4, 5}
	fmt.Printf("The length of the vector is %f\n", Abs(p2))

	q := Scale(p1, 5)
	fmt.Printf("The length of the vector q is: %f\n", Abs(&q))
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f", q.X, q.Y)

}
