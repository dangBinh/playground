package main

import (
	"fmt"
)

func InsertionSort(sortList []int) []int {
	for i := 1; i < len(sortList); i++ {
		currentValue := sortList[i]
		position := i
		for position > 0 && sortList[position-1] > currentValue {
			sortList[position] = sortList[position-1]
			position--
		}
		sortList[position] = currentValue
	}
	return sortList
}

func main() {
	myList := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println(InsertionSort(myList))
}
