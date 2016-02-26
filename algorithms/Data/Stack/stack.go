package main

import (
	"fmt"
)

type Stack struct {
	s []interface{}
}

func NewStack(number uint) *Stack {
	return &Stack{s: make([]interface{}, 0, number)}
}

func (s *Stack) Push(item interface{}) {
	s.s = append(s.s, item)
}

func (s *Stack) Pop() interface{} {
	l := len(s.s)
	if l > 0 {
		res := s.s[l-1]
		s.s = s.s[:l-1]
		return res
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	l := len(s.s)
	return s.s[l-1]
}

func (s *Stack) Size() int {
	return len(s.s)
}

func main() {
	myStack := NewStack(10)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push("Hello")
	fmt.Println(myStack)
}
