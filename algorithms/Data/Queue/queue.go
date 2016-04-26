package main

import (
	"fmt"
)

type Queue struct {
	q []interface{}
}

func NewQueue(number uint) *Queue {
	return &Queue{q: make([]interface{}, 0, number)}
}

func (q *Queue) Enqueue(item interface{}) {
	temp := make([]interface{}, 0, 1)
	temp = append(temp, item)
	q.q = append(temp, q.q...)
}

func (q *Queue) Dequeue() interface{} {
	l := len(q.q)
	if l > 0 {
		res := q.q[l-1]
		q.q = q.q[:l-1]
		return res
	}
	return nil
}

func (q Queue) Size() int {
	return len(q.q)
}

func main() {
	queue := NewQueue(10)
	queue.Enqueue(10)
	fmt.Println(queue.Size())
	fmt.Println(queue.Dequeue())
}
