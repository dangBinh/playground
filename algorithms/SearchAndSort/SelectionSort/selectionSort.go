package main

import (
	"fmt"
)

func SelectionSort(sortList []int) []int {
	for slot := len(sortList) - 1; slot >= 0; slot-- {
		positionOfMax := 0
		for location := 1; location < slot+1; location++ {
			if sortList[location] > sortList[positionOfMax] {
				positionOfMax = location
			}
		}
		sortList[slot], sortList[positionOfMax] = sortList[positionOfMax], sortList[slot]
	}
	return sortList
}

func main() {
	mylist := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println(SelectionSort(mylist))
}
