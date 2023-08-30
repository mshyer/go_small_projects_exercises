package main

import (
	"fmt"
)

/*PEDAC
input: array of chars (bytes)
output: mutate the same array with O(1) space
Problem: reverse the chars
ALGO TYPE METHOD: two pointer
move the pointers i, j until they meet, swapping chars


*/

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
	fmt.Printf("%s\n", s1)

}