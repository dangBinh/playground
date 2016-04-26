package main

import (
	"fmt"
)

func BinarySearch(myList []int, item int) bool {
	first := 0
	last := len(myList) - 1
	found := false
	for !found && first <= last {
		middle := (last + first) / 2
		if myList[middle] == item {
			found = true
		} else {
			if item < myList[middle] {
				last = myList[middle] - 1
			} else {
				first = myList[middle] + 1
			}
		}
	}
	return found
}

func main() {
	myList := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(BinarySearch(myList, 7))
}
