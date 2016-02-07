package main

import (
	"fmt"
)

func main() {
	var age, height, weight int
	var s string
	var i int
	fmt.Println("How old are you?")
	fmt.Scanf("%d", &age)
	fmt.Println("How tall are you?")
	fmt.Scanf("%d", &height)
	fmt.Println("How much do you weight?")
	fmt.Scanf("%d", &weight)

	fmt.Printf("%d, %d, %d", age, height, weight)

	fmt.Sscanf(" 12 34 567", "%5s%d", &s, &i)

	fmt.Printf("%v", s)
	fmt.Printf("%v", i)
}
