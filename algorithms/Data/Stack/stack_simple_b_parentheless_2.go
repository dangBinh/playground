package main

import (
	"fmt"
	"strings"
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

func ParChecker(symbols string) bool {
	var symbolsLength = uint(len(symbols))
	s := NewStack(symbolsLength)
	balanced := true
	pattern := "({["
	i := 0
	for i < len(symbols) && balanced {
		symbol := string(symbols[i])
		if strings.Index(pattern, symbol) != -1 {
			s.Push(symbol)
		} else {
			if s.IsEmpty() {
				balanced = false
			} else {
				top := s.Pop().(string)
				if !Matches(top, symbol) {
					balanced = false
				}
			}
		}
		i++
	}
	if s.IsEmpty() && balanced {
		return true
	} else {
		return false
	}
}

func Matches(open string, close string) bool {
	pattern_open := "({["
	pattern_close := ")}]"
	return strings.Index(pattern_open, open) == strings.Index(pattern_close, close)
}

func main() {
	fmt.Println(ParChecker("(({{}}))"))
	fmt.Println(ParChecker("(()"))
}
