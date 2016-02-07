package main

import (
	"fmt"
)

func main() {
	str1 := "abcd"
	str2 := "bcda"

	fmt.Println(anagramSolution1(str1, str2))
}

func anagramSolution1(str1 string, str2 string) bool {
	ok := true
	pos1 := 0
	for pos1 < len(str1) && ok {
		pos2 := 0
		found := false
		for pos2 < len(str2) && !found {
			if str1[pos1] == str2[pos2] {
				found = true
			} else {
				pos2 += 1
			}
		}

		if found {
			str2 = str2[:pos2] + str2[pos2+1:]
		} else {
			ok = false
		}

		pos1 += 1
	}

	return ok
}
