package main

import (
	"fmt"
)

/*PEDAC: array pre-sorted
input: pre-sorted array of ints and target int
output: array of 2 ints that add to target. 
	CONSTRAINTS: Exactly one solution
pointer method: START-END POINTER NOTE: works only with sorted array
1: i init 0
2: j init len - 1
3: loop until i and j meet
  -if j + i == target, return [i + 1, j + 1] (index on leetcode starts from 1)
	-If j + i < target, iterate i
	-if j + i > target, iterate j--

*/

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func main() {
	myArr := []int{2, 7, 11, 15}
	myArr2 := []int{-1, 0}
	myArr3 := []int{2, 3, 4}
	fmt.Println(twoSum(myArr, 9))
	fmt.Println(twoSum(myArr2, -1))
	fmt.Println(twoSum(myArr3, 6))
}