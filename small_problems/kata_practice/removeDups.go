package main

import (
	"fmt"
)

/*PEDAC
input: array of non-decreasing integers
outputs:
  1. MUTATE input array such that k num unique elements appear first
    -subsequent elements replaced with _
  2. RETURN num unique elements
ALGORITHM TYPE: anchor/runner
init counter "uniqueValues" to 1
MOVE RUNNER: always
MOVE ANCHOR: before every swap. Increment counter with every swap.
SWAP CONDITION: when nums(runner) != nums(anchor), move anchor up 1 and swap.
CLEANUP
move anchor to uniqueValues counter, move to end of array, replacing with _
*/


func removeDuplicates(nums []int) int {
	//init anchor, runner, unique value counter
	a, r, uniqueValues := 0, 1, 1

	//increment the runner, swapping if runner doesn't equal anchor
	for r < len(nums) {
		if nums[r] != nums[a] {
			a++
			uniqueValues++
			nums[a] = nums[r]
		}
		r++
	}
	//cleanup
	for a < len(nums) - 1{
		a++
		nums[a] = '_'
	}
	return uniqueValues
}

func main() {
	arr1 := []int{1, 2, 2} // => [1, 2, _] && 2
	arr2 := []int{0,0,1,1,1,2,2,3,3,4} //=> [0, 1, 2, 3, 4, _ _ _ _ _] && 5
	fmt.Println(removeDuplicates(arr1))
	fmt.Println(arr1)
	fmt.Println(removeDuplicates(arr2))
	fmt.Println(arr2)
}