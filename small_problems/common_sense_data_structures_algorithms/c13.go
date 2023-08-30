package main

import (
	"fmt"
	"sort"
)

type SortableSlice struct {
	values []int
}

func (s *SortableSlice) partition(lPoint, rPoint int) int {
	pivotIdx := rPoint
	pivot := s.values[pivotIdx]
	rPoint -= 1

	for true {

		for s.values[lPoint] < pivot {
			lPoint += 1
		}

		for s.values[rPoint] > pivot {
			rPoint -= 1
		}

		if lPoint >= rPoint {
			break
		} else {
			rVal := s.values[rPoint]
			lVal := s.values[lPoint]
			s.values[lPoint] = rVal
			s.values[rPoint] = lVal

			lPoint += 1
		}
	}
	lVal := s.values[lPoint]
	s.values[pivotIdx] = lVal
	s.values[lPoint] = pivot

	return lPoint
}

func (s *SortableSlice) QuickSort(lIdx, rIdx int) {
	if rIdx - lIdx <= 0 {
		return
	}
	pivotIdx := s.partition(lIdx, rIdx)

	s.QuickSort(lIdx, pivotIdx - 1)
	s.QuickSort(pivotIdx, rIdx)
}

func (s *SortableSlice) QuickSelect(kthLowestValue, lIdx, rIdx int) int {
	if rIdx - lIdx <= 0 {
		return s.values[lIdx]
	}

	pivotIdx := s.partition(lIdx, rIdx)

	if kthLowestValue < pivotIdx {
		return s.QuickSelect(kthLowestValue, lIdx, pivotIdx - 1)
	} else if kthLowestValue > pivotIdx {
		return s.QuickSelect(kthLowestValue, pivotIdx + 1, rIdx)
	} else {
		return s.values[pivotIdx]
	}
}

/*

input: array of integers
output: integer (greatest product of 3 numbers)
[1, 7, 3, 5, 9, 4, 2, 8] == 9 * 8 * 7
step 1: sort using sort ints. Return product of last 3 sorted

*/
func greatestProductThree(arr []int) int {
	sort.Ints(arr)
	big3 := arr[len(arr) - 3:]
	return big3[0] * big3[1] * big3[2]
}

func missingNo(arr []int) int {
	sort.Ints(arr)
	counter := 0
	for true {
		if counter >= len(arr) {
			return -1
		}

		if arr[counter] != counter {
			return counter
		}
		counter += 1
	}
	return -1
}
/*

input: array of ints
output: int (greatest int)

O(N) -> keep track of the greatest (done already in previous chapter)
O(log n) -> sort array and return last element
O(N^2) => recursive with an extra step
*/
func findBigBoi(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if (arr[0] > findBigBoi(arr[1:])) {
		return arr[0]
	} else {
		return findBigBoi(arr[1:])
	}
}

func main() {
	values := []int{1, 6, 2, 8, 9, 4, 3, 7, 5}
	SSlice := SortableSlice{
		values: values,
	}
	fmt.Println(SSlice.values)
	// fmt.Println(SSlice.partition(0, len(SSlice.values) - 1))
	// SSlice.QuickSort(0, len(SSlice.values) - 1)
	fmt.Println(SSlice.QuickSelect(5, 0, len(SSlice.values) - 1))
	fmt.Println(SSlice.values)

	fmt.Println(greatestProductThree([]int{5, 4, 3, 7, 6, 1, 9, 8, 2}))
	fmt.Println(missingNo([]int{6, 0, 5, 7, 3, 4, 1}))
	fmt.Println(findBigBoi([]int{1, 22, 3, 4}))
}