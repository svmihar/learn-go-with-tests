package main

import "math"

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
