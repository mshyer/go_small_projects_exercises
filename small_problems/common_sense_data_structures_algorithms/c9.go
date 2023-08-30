/*
1: queue: because fairness dictates the first to arrive should get served first
2: 4
3: 3



*/
package main

import (
	"fmt"
)

type strStack struct {
	list []string
}

func (s *strStack) read() string{
	return s.list[len(s.list) - 1]
}

func (s *strStack) push(str string) {
	s.list = append(s.list, str)
}

func (s *strStack) pop() string {
	removed := s.list[len(s.list) - 1]
	s.list = s.list[:len(s.list) - 1]
	return removed
}

func (s *strStack) length() int {
	return len(s.list)
}

func reverseString(str string) string {
	myStack := new(strStack)
	for _, chr := range str {
		myStack.push(string(chr))
	}
	var outputStr string
	length := myStack.length()
	fmt.Println(length)
	for idx := 0; idx < length; idx++ {
		outputStr += myStack.pop()
	}
	return outputStr
}

func main() {
	fmt.Println(reverseString("abcde"))
}