package main

import (
	"fmt"
)

func main() {
	s1 := "abcd"
	s2 := "bcda"

	fmt.Println(AnagramSolution3(s1, s2))
}

func AnagramSolution3(s1 string, s2 string) bool {
	var c1 [26]int
	var c2 [26]int

	for _, w1 := range s1 {
		pos := int(w1) - int('a')
		c1[pos] = 1
	}

	for _, w2 := range s2 {
		pos := int(w2) - int('a')
		c2[pos] = 1
	}

	j := 0
	ok := true
	for j < 26 && ok {
		if c1[j] == c2[j] {
			j++
		} else {
			ok = false
		}
	}

	return ok
}
