package main

import (
	"fmt"
	"math"
)

func ReverseIt(arr []int) {
	//[1, 2, 3, 4, 5] output [5, 4, 3, 2, 1] same array
	//1 2 3 5 4
	//1 2 5 4 3
	//Start n at array length - 1
	//push each item to end of array
	  //1 2 3 5 4 n == 2
		  //for i = n i <= arr.lengh - 2 i ++ 
				//arr [i] = arr [i + 1] and vice versa
	// for n := len(arr); n >= 0; n-- {
	// 	for idx := n; idx <= len(arr) - 2; idx++ {
	// 		temp := arr[idx + 1]
	// 		arr[idx+ 1] = arr[idx]
	// 		arr[idx] = temp
	// 	} 
	// }
	// return

	// n = 0 n < arr.length / 2 n++
	//temp := arr[n]
	//arr[len(arr) - n]
	for n := 0; n < (len(arr) / 2); n++ {
		temp := arr[n]
		arr[n] = arr[len(arr) - n - 1]
		arr[len(arr) - n - 1] = temp
	}
	return

}
// 
func gameWinner(numCoins int, currentPlayer bool) bool {
	if numCoins <= 0 {
		return currentPlayer
	}
	if gameWinner(numCoins - 1, !currentPlayer) == currentPlayer || gameWinner(numCoins - 2, !currentPlayer) == currentPlayer {
		return currentPlayer
	} else {
		return !currentPlayer
	}
}
func SumSwap(arr1 *[]int, arr2 *[]int) {
	//function takes two arrays
	//input: two arrays of numbers
	//Output: none, arrays mutated or not mutated
	//Find the difference between sums (22 - 8 = 14)
	//Find if there is a set of numbers where in the larger array, it is larger
	   //than the number in the smaller array by half the difference. IE 9 2 (7)
		 	//Nested for loops: for n in large array
			  //for j in smaller array
				   //if n - j = arrayDifference / 2
					 //arr1[n] = arr2[j] (swap so use temp)
		 //Then swap the two numbers
	
	(*arr1)[3] = 10
}

func main() {
	myArr := []int{1, 2, 3, 4, 5}
	fmt.Println(myArr)
	ReverseIt(myArr)
	fmt.Println(myArr)
	fmt.Println(gameWinner(7, true))
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	SumSwap(&arr1, &arr2)
	fmt.Println(arr1)
	fmt.Println()
}