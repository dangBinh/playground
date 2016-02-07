package main

import (
	"fmt"
)

func main() {
	cars := 90
	spaces_in_a_car := 4.0
	drivers := 30
	passengers := 90
	cars_not_driven := cars - passengers
	cars_driven := drivers
	carpool_capacity := cars * int(spaces_in_a_car)
	average_passergers_per_car := passengers / cars_driven

	fmt.Printf("There are %d cars available \n", cars)
	fmt.Printf("There are only %d drivers available \n", drivers)
	fmt.Printf("There will be %d empty cars to day \n", cars_not_driven)
	fmt.Printf("We can transport %d people today \n", carpool_capacity)
	fmt.Printf("We have %d to carpool today \n", passengers)
	fmt.Printf("We need to put about %d in each car \n", average_passergers_per_car)
}
