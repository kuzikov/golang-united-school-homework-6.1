package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rr *Rectangle) CalcPerimeter() float64 {
	return (rr.Height + rr.Weight) * 2 // P = (a + b) * 2
}

func (rr *Rectangle) CalcArea() float64 {
	return rr.Height * rr.Weight // S = a * b
}
