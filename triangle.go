package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (tt *Triangle) CalcPerimeter() float64 {
	return 3 * tt.Side // P = 3 * a
}

func (tt *Triangle) CalcArea() float64 {
	return (math.Sqrt(3) / 4) * math.Pow(tt.Side, 2) // S = (sqrt(3) / 4) * a^2
}
