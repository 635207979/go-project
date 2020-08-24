package Graphics

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Hight float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Hight
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}

func Perimeter(r Rectangle) float64 {
	result := (r.Width + r.Hight) * 2
	return result
}
