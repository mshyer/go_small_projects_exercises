package main

import (
	"fmt"
)

type Result struct {
	C rune
	L int
}

func LongestRepetition(text string) Result {
	//two pointer strategy
	//return if empty

	var maxChar rune
	var maxCount int = 0
	//set i, j to 0
	//set maxCount to 0
	//while j < text length
	// while j < text length && textj == texti {
	//j Plus one
	//if j - i > maxCount
	//set maxChar to index of i
	//set maxCount to j - i
	//set i to j
	i := 0
	j := 0
	greedyResult := new(Result)
	for j < len(text) {
		for j < len(text) && (text[j] == text[i]) {
			j++
		}
		if j-i > maxCount {
			maxChar = rune(text[i])
			maxCount = j - i
		}
		i = j
	}
	greedyResult.C = maxChar
	greedyResult.L = maxCount
	return *greedyResult
}

func main() {
	a1 := LongestRepetition("aaaabb")
	a2 := LongestRepetition("bbbaaabaaaa")
	a3 := LongestRepetition("cbdeuuu900")
	a4 := LongestRepetition("abbbbb")
	a5 := LongestRepetition("aabb")
	a6 := LongestRepetition("ba")
	a7 := LongestRepetition("")
	fmt.Println(a1, a2, a3, a4, a5, a6, a7)
}
