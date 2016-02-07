package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "bcda"

	fmt.Println(AnagramSolution2(s1, s2))
}

func AnagramSolution2(s1 string, s2 string) bool {
	s1 = SortString(s1)
	s2 = SortString(s2)

	pos1 := 0
	ok := true

	for pos1 < len(s1) && ok {
		if s1[pos1] == s2[pos1] {
			pos1 += 1
		} else {
			ok = false
		}
	}

	return ok
}

func SortString(s string) string {
	c := strings.Split(s, "")
	sort.Strings(c)
	return strings.Join(c, "")
}
