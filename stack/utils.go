package stack

import "errors"

type Stack[E any] interface {
	Push(value E)
	Pop() (E, error)
	Peak() (E, error)
	Size() int
}

type ArrayStack[E any] []E

func (s *ArrayStack[E]) Push(value E) {
	*s = append(*s, value)
}

func (s *ArrayStack[E]) Pop() (E, error) {
	if len(*s) == 0 {
		var zero E
		return zero, errors.New("invalid pop on empty stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res, nil
}

func (s *ArrayStack[E]) Peak() (E, error) {
	if len(*s) == 0 {
		var zero E
		return zero, errors.New("invalid peak on empty stack")
	}

	return (*s)[len(*s)], nil
}

func (s *ArrayStack[E]) Size() int {
	return len(*s)
}
