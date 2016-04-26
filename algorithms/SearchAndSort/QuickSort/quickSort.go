package main

import (
	"fmt"
)

func QuickSort(sortList []int) []int {
	QuickSortHelper(sortList, 0, len(sortList)-1)
	return sortList
}

func QuickSortHelper(sortList []int, first int, last int) {
	if first < last {
		splitPoint := Partition(sortList, first, last)

		QuickSortHelper(sortList, first, splitPoint-1)
		QuickSortHelper(sortList, splitPoint+1, last)
	}
}

func Partition(sortList []int, first int, last int) int {
	pivotValue := sortList[first]
	leftMark := first + 1
	rightMark := last
	found := false
	for !found {
		for leftMark <= rightMark && sortList[leftMark] <= pivotValue {
			leftMark++
		}

		for rightMark >= leftMark && sortList[rightMark] >= pivotValue {
			rightMark--
		}
		if rightMark < leftMark {
			found = true
		} else {
			sortList[leftMark], sortList[rightMark] = sortList[rightMark], sortList[leftMark]
		}
	}
	sortList[first], sortList[rightMark] = sortList[rightMark], sortList[first]
	return rightMark
}

func main() {
	myList := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println(QuickSort(myList))
}
