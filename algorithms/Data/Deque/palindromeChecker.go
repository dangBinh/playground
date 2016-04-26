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
	res := d.d[0]
	d.d = d.d[1:]
	return res
}

func (d *Deque) Size() int {
	return len(d.d)
}

func PalChekcer(str string) bool {
	strLen := uint(len(str))
	deque := NewDeque(strLen)

	// Add to deque
	for _, ch := range str {
		deque.AddRear(ch)
	}

	stillEqual := true

	for deque.Size() > 1 && stillEqual {
		first := deque.RemoveFront()
		last := deque.RemoveRear()
		if first != last {
			stillEqual = false
		}
	}
	return stillEqual
}

func main() {
	fmt.Println(PalChekcer("radar"))
}
