package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type UnorderedList struct {
	head *Node
}

func (ul *UnorderedList) First() *Node {
	return ul.head
}

func (ul *Node) Next() *Node {
	return ul.next
}

func (ul *UnorderedList) Add(value int) *UnorderedList {
	node := &Node{value: value}
	if ul.head == nil { // First node
		ul.head = node
	} else {
		temp := node
		temp.next = ul.head
		ul.head = temp
	}
	return ul
}

func (ul *UnorderedList) Search(value int) bool {
	current := ul.head
	found := false
	for current != nil && !found {
		if current.value == value {
			found = true
		} else {
			current = current.next
		}
	}
	return found
}

func (ul *UnorderedList) Size() int {
	current := ul.head
	size := 0
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (ul *UnorderedList) Index(value int) int {
	current := ul.head
	found := false
	index := 0
	for current != nil && !found {
		if current.value == value {
			found = true
		} else {
			index++
			current = current.next
		}
	}
	return index
}

func (ul *UnorderedList) Remove(value int) {
	current := ul.head
	previous := &Node{}
	found := false
	for !found {
		if current.value == value {
			found = true
		} else {
			previous = current
			current = current.next
		}
	}

	if previous.value == 0 {
		ul.head = current.next
	} else {
		previous.next = current.next
	}
}

func (ul *UnorderedList) Insert(pos int, value int) *UnorderedList {
	current := ul.head
	previous := &Node{}
	found := false
	index := 0
	for current != nil && !found {
		if pos == index {
			found = true
		} else {
			index++
			previous = current
			current = current.next
		}
	}
	if previous.value == 0 {
		ul.Add(value)
	} else if current != nil {
		temp := &Node{value: value, next: current}
		previous.next = temp
	}
	return ul
}

func (ul *UnorderedList) Pop(pos int) *UnorderedList {
	current := ul.head
	previous := &Node{}
	found := false
	index := 0
	for current != nil && !found {
		if pos == index {
			found = true
		} else {
			index++
			previous = current
			current = current.next
		}
	}

	if previous.value == 0 { // Remove first node
		ul.head = current.next
	} else {
		previous.next = current.next
	}

	return ul
}

func main() {
	l := new(UnorderedList)
	l.Add(1)
	l.Add(2)
	fmt.Println(l)
	fmt.Println(l.Search(2))
	fmt.Println(l.Size())
	fmt.Println(l.Index(1))
	l.Remove(2)
	fmt.Println(l.Search(2))
	l.Add(2)
	l.Insert(1, 3)
	fmt.Println(l.Search(3))
	l.Pop(1)
	fmt.Println(l.Search(3))
}
