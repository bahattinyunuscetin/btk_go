package interface_

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func School(s shape) {
	fmt.Println("şeklin alani:", s.area())
}
func Demo1() {
	r := rectangle{width: 10, height: 6}
	School(r)
	c := circle{radius: 10}
	School(c)
}
