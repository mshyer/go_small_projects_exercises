package main

/*
1: write a function that retuns an intersection of two arrays
input: two arrays of integers
output: array of integers
Restriction: algorithm must be O(N)



*/

import (
	"fmt"
	"strings"
)

var pri = fmt.Println

func intersection(arr1, arr2 []int) (intersection []int) {
	//1: find the larger of the two arrays
	//2: convert the larger array into a hash
	//3: look up each item in smaller array
	//4: for each item in the hash, push it to the output array
	//5: return the output array
	if len(arr2) > len(arr1) {
		placeholder := arr1
		arr1 = arr2
		arr2 = placeholder
	}
	hashTable := make(map[int]bool)
	for idx := 0; idx < len(arr1); idx++ {
		hashTable[arr1[idx]] = true
	}

	for idx := 0; idx < len(arr2); idx++ {
		if hashTable[arr2[idx]] {
			intersection = append(intersection, arr2[idx])
		}
	}
	return
}


func findDup(strArr []string) string {
	//put the strings into a map with true as you iterate
	// if the string is already in the map, return that string
	hashTable := make(map[string]bool)
	for idx := 0; idx < len(strArr); idx++ {
		if hashTable[strArr[idx]] {
			return strArr[idx]
		} else {
			hashTable[strArr[idx]] = true
		}
	}

	return ""
}

func missingWhichLetter(str string) rune {
	//add all of the letters in alphabet to a hash (26)
	//for each char in string, switch the alphabet to off (N)
	//for each letter in the alphabet, find the ONE that's off (26) N + X ==O(N)
	alphabetHash := make(map[rune]bool)
	str = strings.ToLower(str)


	//for each char in string, set the hash to false: N steps
	for idx := 0; idx < len(str); idx++ {
		alphabetHash[rune(str[idx])] = true
	}

	//Find the true value in alphabet hash: 26 steps

	for idx := 97; idx < 123; idx++ {
		if !alphabetHash[rune(idx)] {
			return rune(idx)
		}
	}
	return rune(0)
}

func findFirstNonDup(str string) rune {
	//Loop through once and add to hash plus count
	//Loop through again, and if the count isn't 2, return that (2n)

	hashTable := make(map[rune]int)
	for _, chr := range str {
		hashTable[chr]++
	}

	for _, chr := range str {
		if hashTable[chr] < 2 {
			return chr
		}
	}
	return rune(0)
}



func main() {
	var arr1 = []int{1, 2, 3, 4, 44}
	var arr2 = []int{5, 6, 7, 8, 9, 0, 44, 55, 33, 66, 2, 3}
	pri(intersection(arr1, arr2))


	strArr1 := []string{"happy", "as", "a", "clam", "clam", "chowder"}
	pri(findDup(strArr1))
	pri(string(97))
	pri(string(122))

	pri(string(missingWhichLetter("the quick brown box jumps over a lazy dog")))
	pri(string(findFirstNonDup("minimum")))
}