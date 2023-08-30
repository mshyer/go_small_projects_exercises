package main

import (
	"fmt"
	"strings"
)

// function convers dash-underscore to CamelCase
// words always capitalized except first word.
// First word capitalized if original first word capitalized
// Input: string
// Output: CamelCaseString
func ToCamelCase(s string) string {
	//Two pointer algorithm:
	//J points to the start of a word
	//1: firstLetter := find the first letter and store it
	//loop until not "-" or "_"
	//Store that in variable
	//BIG WHILE: while j < string length
	//Move j until you hit an underscore or dash
	//While s[j] !== "_" || "-"
	//Move i until it reaches j
	//If index 0, push Capitalized to the output string
	//Else push lowercase to the output
	//set i == j
	//move j until you hit another letter
	//while j == "_" || "-"
	//set i = j
	//set the first letter to equal firstLetter
	//return the string
	var firstLetter rune
	firstIdx := 0
	output := ""
	for idx, char := range s {
		if char != '-' && char != '_' {
			firstLetter = char
			firstIdx = idx
			break
		}
	}
	i := firstIdx
	j := firstIdx
	for j < len(s)-1 {
		for j < len(s) && s[j] != '-' && s[j] != '_' {
			j += 1
		}
		output += strings.ToUpper(string(s[i]))
		i++
		for i < j {
			output += strings.ToLower(string(s[i]))
			i++
		}
		for j < len(s) && (s[j] == '_' || s[j] == '-') {
			j++
		}
		i = j
	}
	output = string(firstLetter) + output[1:]
	return output
}

func main() {
	a1 := ToCamelCase("the-stealth-warrior")
	fmt.Println(a1)
	a2 := ToCamelCase("The_Stealth_Warrior")
	fmt.Println(a2)
	a3 := ToCamelCase("The_Stealth-Warrior")
	fmt.Println(a3)
	a4 := ToCamelCase("TheStealthWarrior")
	fmt.Println(a4)
	a5 := ToCamelCase("_-TheStealthWarrior")
	fmt.Println(a5)

}
