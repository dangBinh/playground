package main

import (
	"fmt"
)

func sum(numList []int) int {
	if len(numList) == 1 {
		return numList[0]
	} else {
		return numList[0] + sum(numList[1:])
	}
}

func main() {
	numList := []int{1, 2, 3}
	fmt.Println(sum(numList))
}
