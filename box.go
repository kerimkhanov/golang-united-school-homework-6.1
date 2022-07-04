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
	if b.shapesCapacity == len(b.shapes) {
		e := fmt.Sprintf("addition limit exceeded, shapes capacity is %d", b.shapesCapacity)
		return errors.New(e)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	// if i <0 || i>=len(b.shapes)	{
	// 	return nil, errors.New("out of range")
	// }
	if i > b.shapesCapacity-1 {
		e := fmt.Sprintf("index is out of capasity, limit is %d", b.shapesCapacity)
		return nil, errors.New(e)
	}
	if i < b.shapesCapacity-1 && i > len(b.shapes)-1 {
		e := fmt.Sprintf("shape with index %d not founded", i)
		return nil, errors.New(e)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	_, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	res := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return res, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("can not replace shape with index %d: %s", i, err.Error())
	}
	s := b.shapes[i]
	b.shapes[i] = shape
	return s, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, w := range b.shapes {
		sum += w.CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, w := range b.shapes {
		sum += w.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	// for _, w := range b.shapes {
	// 	if w == b.shapes.(*Circle) {

	// 	}
	// }
	count := 0
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			b.ExtractByIndex(i)
			i--
			count++
		} else if _, ok := b.shapes[i].(Circle); ok {
			b.ExtractByIndex(i)
			i--
			count++
		}

	}
	if count == 0 {
		return fmt.Errorf("there is no circles")
	}
	return nil
}
