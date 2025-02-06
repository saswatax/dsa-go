package stack

import "errors"

type Stack interface {
	Push(value int)
	Pop() (int, error)
	Peak() (int, error)
	Size() int
}

type ArrayStack[S ~[]E, E any] []S

func (s *ArrayStack[any]) Push(value int) {
	*s = append(*s, value)
}

func (s *ArrayStack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("invalid pop on empty stack")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res, nil
}

func (s *ArrayStack) Peak() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("invalid peak on empty stack")
	}

	return (*s)[len(*s)], nil
}

func (s *ArrayStack) Size() int {
	return len(*s)
}
