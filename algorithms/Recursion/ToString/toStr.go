package main

import (
	"fmt"
)

func toStr(number int, base int) string {
	convertString := "0123456789ABCDEF"
	if number < base {
		return string(convertString[number])
	} else {
		return toStr(number/base, base) + string(convertString[number%base])
	}
}

func main() {
	fmt.Println(toStr(1435, 16))
}
