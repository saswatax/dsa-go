package stack

import "errors"

type Stack[E any] interface {
	Push(value E)
	Pop() (E, error)
	Peak() (E, error)
	Size() int
}

type ArrayStack[E any] struct {
	arr []E
}

func (s *ArrayStack[E]) Push(value E) {
	s.arr = append(s.arr, value)
}

func (s *ArrayStack[E]) Pop() (E, error) {
	top := len(s.arr) - 1
	if top == -1 {
		var zero E
		return zero, errors.New("invalid pop on empty stack")
	}

	res := s.arr[top]
	s.arr = s.arr[:top]

	return res, nil
}

func (s *ArrayStack[E]) Peak() (E, error) {
	top := len(s.arr) - 1
	if top == -1 {
		var zero E
		return zero, errors.New("invalid peak on empty stack")
	}

	return s.arr[top], nil
}

func (s *ArrayStack[E]) Size() int {
	return len(s.arr)
}
