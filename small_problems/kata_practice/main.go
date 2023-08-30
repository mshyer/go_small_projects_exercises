package main

import (
	"fmt"
)

func main() {
	var s, t int
	fmt.Scan(&s, &t)
	fmt.Println(s, t)

	maxDistance := 0
	baseSpeed := s

	for i := 0; i < t; i++ {
		if i%3 == 2 {
			// sprint
			distance := (baseSpeed * 2)
			maxDistance += distance
			baseSpeed--
		} else {
			// normal run or recovery
			distance := baseSpeed
			maxDistance += distance
		}
	}

	fmt.Println(maxDistance)
}
