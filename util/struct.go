package util

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) SetAttribute(width float64, height float64) {
	r.width = width
	r.height = height
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
