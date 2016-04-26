package main

import (
	"fmt"
)

func InArray(list []int, check int) bool {
	for _, value := range list {
		if value == check {
			return true
		}
	}
	return false
}

func Filter(list []int, f func(v int) bool) []int {
	listFilter := make([]int, 0)
	for _, value := range list {
		if f(value) {
			listFilter = append(listFilter, value)
		}
	}
	return listFilter
}

func RecDC(coinValueList []int, change int, knowResult []int) int {
	minCoin := change
	if InArray(coinValueList, change) {
		return 1
	} else if knowResult[change] > 0 {
		return knowResult[change]
	} else {
		coinValueList = Filter(coinValueList, func(value int) bool {
			if value <= change {
				return true
			}
			return false
		})
		for _, value := range coinValueList {
			numCoins := 1 + RecDC(coinValueList, change-value, knowResult)
			if numCoins < minCoin {
				minCoin = numCoins
				knowResult[change] = minCoin
			}
		}
	}
	return minCoin
}

func main() {
	list := []int{1, 5, 10, 25}
	knowResult := make([]int, 65)
	fmt.Println(knowResult[63])
	fmt.Println(RecDC(list, 63, knowResult))
}
