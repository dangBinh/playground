package main

import (
	"fmt"
)

func sequentialSearch(myList []int, item int) bool {
	found := false
	pos := 0
	for pos < len(myList) && !found {
		if myList[pos] == item {
			found = true
		} else {
			pos++
		}
	}
	return found
}

func main() {
	myList := []int{1, 2, 3, 4}
	fmt.Println(sequentialSearch(myList, 4))
}
