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

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Peek() interface{} {
	l := len(s.s)
	return s.s[l-1]
}

func (s *Stack) Size() int {
	return len(s.s)
}

func BaseConverter(decimalNumber int, base int) string {
	s := NewStack(10)
	remain := 0
	digits := "0123456789ABCDEF"
	for decimalNumber > 0 {
		remain = decimalNumber % 2
		s.Push(remain)
		decimalNumber = decimalNumber / 2
	}
	var binary string
	for !s.IsEmpty() {
		top := string(digits[s.Pop().(int)])
		binary += top
	}

	return binary
}

func main() {
	fmt.Println(BaseConverter(25, 2))
}
