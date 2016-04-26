package main

import (
	"fmt"
)

func MergeSort(sortList []int) []int {
	if len(sortList) <= 1 {
		return sortList
	}

	middle := len(sortList) / 2

	//note how we slice the array using [:middle]
	left := MergeSort(sortList[:middle])
	//next slice middle to end of array
	right := MergeSort(sortList[middle:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	for i := 0; len(left) > 0 || len(right) > 0; i++ {

		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}

	return result
}

func main() {
	myList := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println(MergeSort(myList))
}
