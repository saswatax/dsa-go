package stack

import "errors"

type Stack interface {
	Push(value int)
	Pop() (int, error)
	Peak() (int, error)
	Size() int
}

type ArrayStack struct {
	arr []int
	top int
}

func (s *ArrayStack) Push(value int) {
	s.arr = append(s.arr, value)
	s.top++
}

func (s *ArrayStack) Pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("invalid pop on empty stack")
	}

	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.top--

	return res, nil
}

func (s *ArrayStack) Peak() (int, error) {
	if s.top == -1 {
		return 0, errors.New("invalid peak on empty stack")
	}

	return s.arr[s.top], nil
}

func (s *ArrayStack) Size() int {
	return s.top + 1
}
