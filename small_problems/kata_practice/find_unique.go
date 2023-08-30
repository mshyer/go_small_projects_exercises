package main

import "fmt"

//INPUT: array with at least 3 float32 numbers
//Arrays can be very large
//OUTPUT: float32, the greatest number in the array
//Strategy: greedy algorithm iterates through array and stores the biggest each loop
//Return biggest at end

//	func FindGreatest(arr []float32) float32 {
//		var greed float32 = arr[0]
//		for idx := 1; idx < len(arr); idx++ {
//			if arr[idx] > greed {
//				greed = arr[idx]
//			}
//		}
//		return greed
//	}
func FindUniq(arr []float32) float32 {
	greed := arr[0]
	if arr[1] != greed {
		if arr[2] == greed {
			return arr[1]
		} else {
			return greed
		}
	}

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] != greed {
			return arr[idx]
		}
	}
	return -1
}

func main() {
	fl1 := FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})
	fmt.Println(fl1)
	fl2 := FindUniq([]float32{0, 0, 0.55, 0, 0})
	fmt.Println(fl2)
	fl3 := FindUniq([]float32{0, 0, 0, 0, 0, 100})
	fmt.Println(fl3)
}
