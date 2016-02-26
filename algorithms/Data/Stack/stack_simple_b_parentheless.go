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

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Peek() interface{} {
	l := len(s.s)
	return s.s[l-1]
}

func (s *Stack) Size() int {
	return len(s.s)
}

func ParChecker(symbols string) bool {
	var symbolsLength = uint(len(symbols))
	s := NewStack(symbolsLength)
	balanced := true
	i := 0
	for i < len(symbols) && balanced {
		symbol := symbols[i]
		if string(symbol) == "(" {
			s.Push(symbol)
		} else {
			if s.isEmpty() {
				balanced = false
			} else {
				s.Pop()
			}
		}
		i++
	}
	if s.isEmpty() && balanced {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(ParChecker("(())"))
	fmt.Println(ParChecker("(()"))
}
