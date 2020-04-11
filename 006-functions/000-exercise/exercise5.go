package main

import "fmt"

type square struct {
	length float32
	width  float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (s square) area() float32 {
	return s.length * s.width
}

func info(s shape) {
	fmt.Println(s.area())
}

func main5() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
		width:  15,
	}

	info(circ)
	info(squa)

}
