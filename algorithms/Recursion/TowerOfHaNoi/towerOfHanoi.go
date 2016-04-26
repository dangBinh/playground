package main

import (
	"fmt"
)

func moveTower(height int, fromPole string, toPole string, withPole string) {
	if height >= 1 {
		moveTower(height-1, fromPole, withPole, toPole)
		moveDisk(fromPole, toPole)
		moveTower(height-1, withPole, toPole, fromPole)
	}
}

func moveDisk(fromPole string, toPole string) {
	fmt.Println("Move from" + fromPole + "to" + toPole)
}

func main() {
	moveTower(3, "A", "B", "C")
}
