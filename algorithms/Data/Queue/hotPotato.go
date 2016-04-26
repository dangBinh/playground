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

func HotPotato(number int, names []string) interface{} {
	nameLen := len(names)
	queue := NewQueue(uint(nameLen))
	// Add name to queue
	for _, name := range names {
		queue.Enqueue(name)
	}

	for queue.Size() > 1 {
		for i := 0; i < number; i++ {
			queue.Enqueue(queue.Dequeue())
		}
		queue.Dequeue()
	}
	return queue.Dequeue()
}

func main() {
	fmt.Println(HotPotato(7, []string{"A", "B", "C", "D", "E", "F"}))
}
