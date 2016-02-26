package main

import (
	"fmt"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) String() string {
	str := ""
	item := s.top
	for {
		if item == nil {
			return str
		}
		if str != "" {
			str = " " + str
		}
		str = "{" + item.value.(string) + "}" + str
		item = item.next
	}
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push("Hello")
	fmt.Sprintf("%v", stack)
}
