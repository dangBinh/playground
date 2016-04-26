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

func (s *Stack) Peek() interface{} {
	l := len(s.s)
	return s.s[l-1]
}

func (s *Stack) Size() int {
	return len(s.s)
}

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

func InfixToPostfix(infixexpr string) string {
	precendece := map[string]int{
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
	}
	opStack := NewStack(uint(len(infixexpr)))
	postfix := make([]string, len(infixexpr))
	tokenList := strings.Split(infixexpr, " ")
	for _, token := range tokenList {
		if strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", token) != -1 || strings.Index("0123456789", token) != -1 {
			postfix = append(postfix, token)
		} else if token == "(" {
			opStack.Push(token)
		} else if token == ")" {
			topToken := opStack.Pop().(string)
			for topToken != "(" {
				postfix = append(postfix, topToken)
				topToken = opStack.Pop().(string)
			}
		} else {
			for !opStack.isEmpty() && precendece[opStack.Peek().(string)] >= precendece[token] {
				postfix = append(postfix, opStack.Pop().(string))
			}
			opStack.Push(token)
		}
	}

	for !opStack.isEmpty() {
		postfix = append(postfix, opStack.Pop().(string))
	}
	return strings.Join(postfix, " ")
}

func main() {
	fmt.Println(InfixToPostfix("A + ( B * C )"))
}
