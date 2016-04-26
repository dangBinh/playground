package main

import (
	"fmt"
)

func BubbleSort(sortList []int) []int {
	for i := len(sortList) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if sortList[j] > sortList[j+1] {
				sortList[j], sortList[j+1] = sortList[j+1], sortList[j]
			}
		}
	}
	return sortList
}

func ShortBubbleSort(sortList []int) []int {
	exchange := true
	sortListLen := len(sortList) - 1
	for sortListLen > 0 && exchange {
		exchange = false
		for i := 0; i < sortListLen; i++ {
			if sortList[i] > sortList[i+1] {
				sortList[i], sortList[i+1] = sortList[i+1], sortList[i]
			}
		}
		sortListLen--
	}
	return sortList
}

func main() {
	myList := []int{1, 20, 10, 5, 2}
	fmt.Println(BubbleSort(myList))
	fmt.Println(ShortBubbleSort(myList))
}
