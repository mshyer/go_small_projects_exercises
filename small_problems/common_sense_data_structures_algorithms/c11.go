package main

import (
	"fmt"
)

var ip = fmt.Println
/*


Heh.

*/

/*
input: array of strings
Output: integer (count num chars in strings)
Data Structure: Recursion

Notes: 
count 1 + count rest
base = len string
*/
func countChars(arr []string) int {
	if len(arr) == 1 {
		return len(arr[0])
	}
	return len(arr[0]) + countChars(arr[1:])
}
/*
input: num array
output: num array (just evens)
data structure: recursion

notes:
[1, 2, 3, 4, 5]
evens 1 + evens rest
end condition: if length 0,
*/

func onlyEvens(arr []int) []int {
	if len(arr) == 1 {
		if arr[0] % 2 == 0 {
			return arr[0:1]
		} else {
			return []int{}
		}
	}
	if arr[0] % 2 == 0 {
		return append(arr[0:1], onlyEvens(arr[1:])...)
	} else {
		return onlyEvens(arr[1:])
	}
}
/*
n + triangle(n - 1)
*/


func triangularNumbers(n int) int {
	if n == 1 {
		return 1
	}
	return n + triangularNumbers(n - 1)
}

/*
PROBLEM: given an arbitrary string w letter x, find index of letter x
input: string
output: int
data structures: recursion

problem: abx  is
  -base case, if str n x return n 
*/

// func findX(str string, arg... int) int {
// 	var idx int
// 	if len(arg) == 0 {
// 		idx = 0
// 	} else {
// 		idx = arg[0]
// 	}
// 	if str[idx] == 'x' || str[idx] == 'X' {
// 		return idx
// 	}
// 	return findX(str, (idx + 1))
// }

func findX(str string) int {
	if str[0] == 'x' {
		return 0
	}
	return findX(str[1:]) + 1
}

/*

input: int, int (n rows) (n columns)
output: int

Notes: simples case: 1 row 1 column  |||

*/
func uniquePaths(cols, rows int) int {
	if cols < 1 || rows < 1 {
		return 0
	} 
	if cols == 1 || rows == 1 {
		return 1
	}

	return uniquePaths(cols - 1, rows) + uniquePaths(cols, rows - 1)
}


func main() {
	strArr := []string{"a", "nice", "young", "Man"}
	ip(countChars(strArr))

	// arr1 := []string{"a", "b", "c"}
	arr2 := []string{"d", "e", "f"}
	ip(append([]string{}, arr2...))
	ip(onlyEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 209}))
	ip(triangularNumbers(6))
	ip(triangularNumbers(7))
	ip(findX("findx"))
	ip(findX("xxxxx"))
	ip(uniquePaths(3, 3))
}