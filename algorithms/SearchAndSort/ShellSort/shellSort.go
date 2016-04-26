package main

import (
	"fmt"
)

func ShellSort(sortList []int) []int {
	countSublist := len(sortList) / 2
	for countSublist > 0 {
		for startPosition := 0; startPosition < countSublist; startPosition++ {
			sortList = GapInsertionSort(sortList, startPosition, countSublist)
			countSublist = countSublist / 2
		}
	}
	return sortList
}

func GapInsertionSort(sortList []int, startPosition int, gap int) []int {
	for i := startPosition + gap; i < len(sortList); i = i + gap {
		currentValue := sortList[i]
		position := i
		for position >= gap && sortList[position-gap] > currentValue {
			sortList[position] = sortList[position-gap]
			position = position - gap
		}
		sortList[position] = currentValue
	}
	return sortList
}

func main() {
	myList := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println(ShellSort(myList))
}
