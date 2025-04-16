package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Square struct {
	Width float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Square) Area() float64 {
	return r.Width * r.Width
}

func printArea(shape Shape) { // as we create method Area so it automatically detects that shape can be rectanle or sqaure
	fmt.Println("Shape is: ", shape.Area())
}
func checkType(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Println("It's a Rectangle with Width:", v.Width)

	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	r := Rectangle{Width: 10, Height: 20}
	sq := Square{Width: 2.2}
	printArea(r)
	printArea(sq)
}
