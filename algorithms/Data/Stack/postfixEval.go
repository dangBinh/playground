package main

import (
	"fmt"
	"strconv"
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

func PostfixEval(postfixexpr string) float64 {
	operandStack := NewStack(uint(len(postfixexpr)))
	tokenList := strings.Split(postfixexpr, " ")

	for _, token := range tokenList {
		if strings.Index("0123456789", token) != -1 {
			if tokenFloat, err := strconv.ParseFloat(token, 64); err == nil {
				operandStack.Push(tokenFloat)
			}
		} else {
			operand1 := operandStack.Pop().(float64)
			operand2 := operandStack.Pop().(float64)
			result := DoMath(token, operand1, operand2)
			operandStack.Push(result)
		}
	}

	return operandStack.Pop().(float64)
}

func DoMath(token string, operand1 float64, operand2 float64) float64 {
	var result float64
	switch token {
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	}
	return result
}

func main() {
	fmt.Println(PostfixEval("8 3 2 + *"))
}
