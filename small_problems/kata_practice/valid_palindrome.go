package main

import (
	"fmt"
	// "regexp"
	"strings"
	"math"
)
var myHash = map[string]bool{
	"a": true,  "b": true, "c": true,
	"d": true,
	"e": true,
	"f": true,
	"g": true,
	"h": true,
	"i": true,
	"j": true,
	"k": true,
	"l": true,
	"m": true,
	"n": true,
	"o": true,
	"p": true,
	"q": true,
	"r": true,
	"s": true,
	"t": true,
	"u": true,
	"v": true,
	"w": true,
	"x": true,
	"y": true,
	"z": true,
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

//Input: string with both alhphanumeric and non-alpha characters
//Output: boolean
//Data Structures: Hash table to remove aphanumerics
//Algorithm: Compare front values to values starting from back (n / 2) -1 steps
//1. Create a hash table that contains all the valid chars
//2. create a new string with only valid chars by using this hash table
//3. lowercase the string
//4. compare front to back

func isCharOrDigit(char byte) bool {
	if myHash[strings.ToLower(string(char))] {
		// fmt.Println("is", string(char))
		return true
	}
	// fmt.Println("not", string(char))
	return false
}

func IsPalindrome(s string) bool {
	// newString := ""
	// for _, byte := range(s) {
	// 	if myHash[strings.ToLower(string(byte))] {
	// 		newString += strings.ToLower(string(byte))
	// 	}
	// }
	// fmt.Println(newString)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isCharOrDigit(s[i]) {
			i++
		}
		for i < j && !isCharOrDigit(s[j]) {
			j--
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true

}

func main() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama")) //true
	fmt.Println(IsPalindrome("race a car")) //false
	fmt.Println(IsPalindrome("")) //true
	fmt.Println(IsPalindrome("A")) //true
	fmt.Println(IsPalindrome(" ")) //true
	fmt.Println(IsPalindrome("*#($&(&#(*&(#*$(*#&$))))) ")) //true
	fmt.Println(math.Abs(1 - 4))
}