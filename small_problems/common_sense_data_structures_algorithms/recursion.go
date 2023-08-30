package main

import (
	"fmt"
)

func factorial(num int) int {
  if num == 1 {
		return 1
	}
	return num * factorial(num - 1)
}

func doubleArray(arr *[]int, idx ...int) (doubled []int) {
	var current int
	if len(idx) == 0 {
		current = 0
	} else {
		current = idx[0]
	}
	(*arr)[current] *= 2
	if current >= len(*arr) - 1 {
		return *arr
	} else {
		return doubleArray(arr, current + 1)
	}
}

func summer(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + summer(arr[1:])
	}
}

func reverseit(str string) string {
	if len(str) == 1 {
		return str
	}
	return reverseit(str[1:]) + string(str[0])
}

func countX(str string) int {
	if len(str) == 1 {
		if str == "x" {
			return 1
		} else {
			return 0
		}
	} else {
		return countX(string(str[0])) + countX(str[1:])
	}
}

func countPaths(n int) int {
	if n < 0 {
		return 0
	} else if n == 1 || n == 0 {
		return 1
	}
	return countPaths(n - 1) + countPaths(n - 2) + countPaths(n - 3)
}

func anagrams(str string) []string {
	//a 
	//anagrams a -> a
	//anagrams ab -> anagrams[a]

	var collection []string
	if len(str) == 1 {
		collection = append(collection, str)
		return collection
	}

	substringAnagrams := anagrams(str[1:])
	fmt.Println(substringAnagrams)

	for _, anagram := range substringAnagrams {
		for idx := 0; idx <= len(anagram); idx++ {
			collection = append(collection, anagram[:idx] + string(str[0]) + anagram[idx:])
		}
	}
	return collection

}



func main() {
	// fmt.Printf("%v\n", factorial(3))
	// fmt.Printf("%v", factorial(6))
	myArr := []int{1, 2, 3, 4}
	doubleArray(&myArr)
	fmt.Printf("%v", myArr)
	fmt.Printf("%v", summer(myArr))

	myString := "avcde"
	fmt.Println(reverseit(myString))

	fmt.Println(countX(	"xbxcxd"))
	// fmt.Println(staircasePaths(5))
	fmt.Println(countPaths(4))
	fmt.Println(anagrams("abcd"))
}