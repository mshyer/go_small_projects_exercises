package main
import (
	"math"
	"fmt"
)
func coordinateValue2(reverse bool, cols, xPos, yPos, matrixAdjust int) int {
	if reverse {
		return (cols * (int(math.Abs(float64(xPos - cols))))) - (cols - (int(math.Abs(float64(xPos - cols)))))
	}
	return (cols * yPos) - (cols - xPos) + matrixAdjust;
}


func spiralOrder2(matrix [][]int) []int {
	var output []int
	cols := len(matrix[0])
	rows := len(matrix)
	totalVals := cols * rows
	matrixAdjust := matrix[0][0] - 1
	reverse := false
	if len(matrix[0]) > 1 {
	  if matrix[0][0] > matrix[0][1] {
		  reverse = true
	  } 
	} else if len(matrix) > 1 {
			if matrix[0][0] > matrix[1][0] {
				reverse = true
			}
	}


	xPos := 1
	yPos := 1
	output = append(output, coordinateValue(reverse, cols, xPos, yPos, matrixAdjust))

	for i := 0; len(output) < totalVals; i++ {
		fmt.Println("big loop")
			for xPos < cols - i {
					xPos++
					output = append(output, coordinateValue(reverse, cols, xPos, yPos, matrixAdjust))
					if len(output) >= totalVals {
							return output
					}
			}
			for yPos < rows - i {
					yPos++
					output = append(output, coordinateValue(reverse, cols, xPos, yPos, matrixAdjust))
					if len(output) >= totalVals {
							return output
					}
			}
			for xPos > 1 + i {
					xPos--
					output = append(output, coordinateValue(reverse, cols, xPos, yPos, matrixAdjust))
					if len(output) >= totalVals {
							return output
					}
			}

			for yPos > 2 + i {
					yPos--
					output = append(output, coordinateValue(reverse, cols, xPos, yPos, matrixAdjust))
					if len(output) >= totalVals {
							return output
					}
			}
	}
	return output
}

func main() {
	fmt.Println(coordinateValue2(true, 1, 1, 1, 0))
	fmt.Println(coordinateValue2(true, 1, 1, 2, 0))
}