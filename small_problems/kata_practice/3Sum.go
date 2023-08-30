package main



import (
	"fmt"
	"sort"
)

/*
input: Integer array
output: nested array of int arrays of length 3 OR empty array
  conditions: sub arrays must be unique integer INDEXES that add to 0
							-solution set must not contain duplicates
	constraints: nums length at least 3

EXAMPLES: [0, 1, 1] => []
					[-1, 0, 1, 2, -1, -4] => [[-1, -1, 2], [-1, 0, 1]]
approach:
	init output
	hash table
	Hash table: approach 1 
	  put all nums in unique hash table plus count
		two loops for i and j through original nums array
		I LOOP until < length n - 2:
		  create a temp hash for the i loop
			J LOOP starts at i plus 1 until end:
		    calculate the companion number
			  if the complement number is in the orig hash table and not 0
			      && not in the temp hash
			    store the j value in the temp hash
				  add the numbers, sorted to output
			increment j by 1
		At end of I loop, decrement the hash value of i by 1

	return the output

*/

func threeSum(nums []int) [][]int {
	// var result [][]int
	companions := make(map[int]int)
	starts := make(map[int]bool)
	output := [][]int{}
	for _, char := range nums {
		companions[char]++
	}

	for i := 0; i < len(nums) - 1; i++ {
		if starts[nums[i]] == true {
			fmt.Println("Starts of nums[i]", i)
			i++
			continue
			fmt.Println("nope")
		}
		tempMap := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			companion := 0 - (nums[j] + nums[i])
			if tempMap[companion] != true && companions[companion] > 0 {
				tempMap[nums[j]] = true
				trio := []int{nums[i], nums[j], companion}
				sort.Ints(trio)
				output = append(output, trio)
			}
			j++
		}
		starts[nums[i]] = true
		
	}


	return output

    
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	// nums := []int{-1,0,1,2,-1,-4}
	// myMap := make(map[int]int)
	// for _, char := range nums {
	// 	myMap[char]++
	// }
	// fmt.Println(myMap)

}