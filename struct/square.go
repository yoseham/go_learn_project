package _struct

import "math"

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radis float64
}
type Triangle struct {
	base float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return (rect.width + rect.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return math.Pi * c.radis * c.radis
}
