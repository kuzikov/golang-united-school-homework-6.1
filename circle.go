package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (cc *Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * cc.Radius // P = 2 * Pi * r
}

func (cc *Circle) CalcArea() float64 {
	return math.Pi * math.Pow(cc.Radius, 2) // S = Pi * r^2
}
