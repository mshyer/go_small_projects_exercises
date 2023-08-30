package main

import "fmt"


/*
input: lowercase strings 
output: bool
thoughts: 
//two pointers
//currentLetter at 0
//r at 0
return true if s is ''
return false if s is not '' and t is ''
for until end of the string
	if runner is next letter, 
		set next letter to next letter in input string
		//if no more characters, return true
		move anchor to runner
	always 
		keep moving the runner
	//if chars left in next return false
	//if no chars left return true

*/
func isSubsequence(s string, t string) bool {
	r := 0
	nextLetter := 0

	if s == "" {
		return true
	}
	for r < len(t) {
		if nextLetter >= len(s) {
			return true
		}
		if t[r] == s[nextLetter] {
			nextLetter++
		}
		r++
	}
	if nextLetter < len(s) {
		return false
	}
	return true
    
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) //true
	fmt.Println(isSubsequence("axc", "ahbgdc")) //false
	fmt.Println(isSubsequence("abc", "abc")) //true
	fmt.Println(isSubsequence("aec", "ahbgdc")) //false
	fmt.Println(isSubsequence("", "")) //true
	fmt.Println(isSubsequence("", "ahbgdc")) //true
	fmt.Println(isSubsequence("b", "abc")) //true
}