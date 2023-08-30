package main

import (
	"fmt"
)


/* PEDAC
input: array with assorted numbers
output: mutate array in place. Zeros should be moved to end
  Requirement: maintain order of non-zero elements
ALGORITHM CATEGORY: Anchor-Runner
1. Start anchor and runner at 0
2. move anchor to first 0
3. move runner to anchor
2. condition to move the runner: runner always moves. Swap with a if non-zero
3. condition to  move the anchor: anchor moves to first 0 position
    //then after swaps, move it up by 1


*/

func moveZeroes(nums []int) {
	a := 0
	r := 0
	for nums[a] != 0 && a < len(nums) - 1{
		a++
	}
	r = a
	for r < len(nums) - 1 {
		r++
		if nums[r] != 0 {
			nums[a], nums[r] = nums[r], nums[a]
			a++
		}
	}
}

func main() {
	myArr := []int{1, 0, 0, 3, 0, 5}
	myArr2 := []int{0, 1, 0, 3, 0, 5}
	myArr3 := []int{6, 1, 9, 5, 3, 6}
	moveZeros(myArr)
	fmt.Println(myArr)
	moveZeros(myArr2)
	fmt.Println(myArr2)
	moveZeros(myArr3)
	fmt.Println(myArr3)

}