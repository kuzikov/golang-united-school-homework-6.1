package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity > len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return errors.New(`capacity exceeded`)
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New(fmt.Sprintf("index %d out of range 0..%d", i, len(b.shapes)-1))
	}

	shape := b.shapes[i]
	if shape == nil {
		return nil, errors.New(`Shape is <nil>`)
	}

	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New(fmt.Sprintf("index %d out of range 0..%d", i, len(b.shapes)-1))
	}

	shape := b.shapes[i] //c
	if shape == nil {
		return nil, errors.New(`Shape is <nil>`)
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New(fmt.Sprintf("index %d out of range 0..%d", i, len(b.shapes)-1))
	}

	removed := b.shapes[i]
	if removed == nil {
		return nil, errors.New(`Shape is <nil>`)
	}

	b.shapes[i] = shape
	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var accum float64 = 0
	for _, sh := range b.shapes {
		accum += sh.CalcPerimeter()
	}

	return accum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var accum float64 = 0
	for _, sh := range b.shapes {
		accum += sh.CalcArea()
	}
	return accum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circleFree := make([]Shape, 0, len(b.shapes))
	eliminated := 0
	for _, sh := range b.shapes {
		if _, ok := sh.(*Circle); ok {
			eliminated++
			continue
		}
		circleFree = append(circleFree, sh)
	}

	if eliminated == 0 {
		return errors.New(`circles are not exists`)
	}
	b.shapes = circleFree
	return nil
}
