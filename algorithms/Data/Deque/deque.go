package main

import (
	"fmt"
)

type Deque struct {
	d []interface{}
}

func NewDeque(number uint) *Deque {
	return &Deque{d: make([]interface{}, 0, number)}
}

func (d *Deque) AddFront(item interface{}) {
	d.d = append(d.d, item)
}

func (d *Deque) AddRear(item interface{}) {
	temp := make([]interface{}, 0, 1)
	temp = append(temp, item)
	d.d = append(temp, d.d...)
}

func (d *Deque) RemoveFront() interface{} {
	res := d.d[len(d.d)-1]
	d.d = d.d[:len(d.d)-1]
	return res
}

func (d *Deque) RemoveRear() interface{} {
	res := d.d[1]
	d.d = d.d[1:]
	return res
}

func main() {
	deque := NewDeque(10)
	deque.AddFront(1)
	deque.AddRear(2)
	fmt.Println(deque)
	deque.RemoveFront()
	fmt.Println(deque)
}
