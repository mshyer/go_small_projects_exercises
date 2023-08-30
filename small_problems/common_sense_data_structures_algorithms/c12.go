package main

import "fmt"

var ft = fmt.Println

func memoFib(n int, memo map[int]int) int {
	memo[1001]+=1
	if n == 0 || n == 1 {
		return 1
	}

	if memo[n] == 0 {
		memo[1000]+= 1
		memo[n] = memoFib(n - 1, memo) + memoFib(n - 2, memo)
	} 
	return memo[n]

}

func fib(n int, memo map[int]int) int {
	memo[1000] += 1
	if n == 0 || n == 1 {
		return 1
	}

	ft("non memo fib", memo[1000])
	return fib(n - 1, memo) + fib(n - 2, memo)
}
/*
input: array of integers
output: integer (sum) less than 100
data structure: recursion
*/

// func addUntil100(arr []int, sum int) int {
// 	if len(arr) == 0 {
// 		return 0
// 	}

// 	var toAdd int
//   if sum + arr[0] > 100 {
// 		toAdd = 0
// 	} else {
// 		toAdd = arr[0]
// 	}

// 	sum += toAdd

// 	return toAdd + addUntil100(arr[1:], sum)
// }

func addUntil100(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sumRemaining := addUntil100(arr[1:])
	if (sumRemaining + arr[0]) > 100 {
		return sumRemaining
	} else {
		return arr[0] + sumRemaining
	}
}

/*

calculate nth number 

*/
func golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if memo[n] == 0 {
		memo[n] = 1 + golomb(n - golomb(golomb(n - 1, memo), memo), memo)
	}

	return memo[n]
}

func uniquePaths2(rows, columns int, memo map[[2]int]int) int {
	// stringsA := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	// str := stringsA[rows] + ":" + stringsA[columns]
	var deets = [2]int{rows, columns}
	if rows == 1 || columns == 1 {
		return 1
	}
	if memo[deets] == 0 {
		memo[deets] = uniquePaths2(rows - 1, columns, memo) + uniquePaths2(rows, columns - 1, memo)
	}
	return memo[deets]
}

func main() {
	ft(memoFib(10, map[int]int{}))
	// ft(fib(10, map[int]int{}))
	ft(addUntil100([]int{1, 2, 3, 8, 100, 1}))
	ft(golomb(4, map[int]int{}))

	ft(fmt.Sprintf("%d", 44))
	ft(uniquePaths2(3, 3, map[[2]int]int{}))
	
}