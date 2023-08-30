package main

import (
	"fmt"
)

func coordinateValue(cols, xPos, yPos int) int {
	return (cols * yPos) - (cols - xPos)
}

func spiralOrder(matrix [][]int) []int {
	var output []int
	cols := len(matrix[0])
	rows := len(matrix)
	totalVals := cols * rows

	xPos := 1
	yPos := 1
	output = append(output, coordinateValue(cols, xPos, yPos))

	for i := 0; len(output) < totalVals; i++ {
		fmt.Println("big loop")
		for xPos < cols-i {
			xPos++
			output = append(output, coordinateValue(cols, xPos, yPos))
			if len(output) >= totalVals {
				return output
			}
		}
		for yPos < rows-i {
			yPos++
			output = append(output, coordinateValue(cols, xPos, yPos))
			if len(output) >= totalVals {
				return output
			}
		}
		for xPos > 1+i {
			xPos--
			output = append(output, coordinateValue(cols, xPos, yPos))
			if len(output) >= totalVals {
				return output
			}
		}

		for yPos > 2+i {
			yPos--
			output = append(output, coordinateValue(cols, xPos, yPos))
			if len(output) >= totalVals {
				return output
			}
		}
	}
	return output
}

func main() {
	// fmt.Println(coordinateValue(2, 1, 1))
	// fmt.Println(coordinateValue(3, 2, 3))
	// fmt.Println(coordinateValue(3, 1, 3))
	// fmt.Println(coordinateValue(3, 2, 2))
	// fmt.Println(coordinateValue(3, 1, 2))
	// fmt.Println(coordinateValue(3, 1, 1))
	m1 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(spiralOrder(m1))
}
