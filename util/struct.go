package util

import "math"

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// SetAttribute ...
func (r *Rectangle) SetAttribute(width float64, height float64) {
	r.Width = width
	r.Height = height
}

// Perimeter ...
func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Area ...
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
