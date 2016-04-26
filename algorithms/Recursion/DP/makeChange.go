package main

import (
	"fmt"
)

func Filter(list []int, f func(v int) bool) []int {
	newList := make([]int, 0)
	for _, value := range list {
		if f(value) {
			newList = append(newList, value)
		}
	}
	return newList
}

func MakeChange(coinValueList []int, change int, minCoins []int, coinUsed []int) int {
	for cents := 0; cents <= change; cents++ {
		coinCount := cents
		newCoin := 1
		coinList := Filter(coinValueList, func(v int) bool {
			if v <= cents {
				return true
			}
			return false
		})
		for _, j := range coinList {
			if minCoins[cents-j]+1 < coinCount {
				coinCount = minCoins[cents-j] + 1
				newCoin = j
			}
		}
		minCoins[cents] = coinCount
		coinUsed[cents] = newCoin
	}
	return minCoins[change]
}

func PrintCoin(coinUsed []int, change int) {
	coin := change
	for coin > 0 {
		thisCoin := coinUsed[coin]
		fmt.Println(thisCoin)
		coin = coin - thisCoin
	}
}

func main() {
	change := 63
	minCoins := make([]int, change+1)
	coinUsed := make([]int, change+1)
	coinValueList := []int{1, 5, 10, 21, 25}
	fmt.Println(MakeChange(coinValueList, change, minCoins, coinUsed))
	PrintCoin(coinUsed, change)
	fmt.Println(coinUsed)
}
