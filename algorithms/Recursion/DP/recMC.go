package main

import (
	"fmt"
)

func Filter(list []int, f func(v int) bool) []int {
	listFilter := make([]int, 0)
	for _, v := range list {
		if f(v) {
			listFilter = append(listFilter, v)
		}
	}
	return listFilter
}

func InArray(list []int, check int) bool {
	for _, v := range list {
		if v == check {
			return true
		}
	}
	return false
}

func RecMC(coinValueList []int, change int) int {
	minCoins := change
	if InArray(coinValueList, change) {
		return 1
	} else {
		coinValueList = Filter(coinValueList, func(value int) bool {
			if value <= change {
				return true
			}
			return false
		})
		for _, coinValue := range coinValueList {
			numCoins := 1 + RecMC(coinValueList, change-coinValue)
			if numCoins < minCoins {
				minCoins = numCoins
			}
		}
	}
	return minCoins
}

func main() {
	list := []int{1, 5, 10, 25}
	fmt.Println(RecMC(list, 23))
}
