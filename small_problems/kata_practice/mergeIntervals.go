package main

import (
	"fmt"
	"sort"
)

/*
input: nested array of 2 ints representing intervals, unsorted
Output: nested array with overlapping invervals merged, in sorted order

0. If 1 interval, return that interval
1: sort the input array
2: create an output array mergedIntervals
3: put the first interval in mergedIntervals
4: iterate through input array and compare to top of merged array
  -if no overlap, append current to mergedArr
	-if partial overlap, set the 1 index of top to the 1 index of current
	- if full overlap, do nothing (drop)
5. return the mergedIntervals
*/


func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sortIntervals(intervals)
	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])

	for _, candidate := range intervals[1:] {
		if top := merged[len(merged)- 1]; candidate[0] > top[1] {
			merged = append(merged, candidate)
		} else if candidate[1] > top[1] {
			top[1] = candidate[1]
		}
	}
	return merged
}

func main() {
}


func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}