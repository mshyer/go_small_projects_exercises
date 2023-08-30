package main

import (
	"fmt"
)

type intStack struct {
	list []int
}

func (s *intStack) push(num int) int {
	s.list = append(s.list, num)
  return len(s.list)
}

func (s *intStack) pop() (val int) {
	val = s.list[len(s.list) - 1]
	s.list = s.list[0:len(s.list) - 1]
	return
}

func (s intStack) read() (val int) {
	return s.list[len(s.list) - 1]
}

func (s *intStack) printStack() {
	for _, val := range s.list {
		fmt.Printf("%d\n", val)
	}
}

func main() {
	myStack := new(intStack)
	myStack.push(3)
	myStack.push(5)
	myStack.push(7)
	myStack.push(11)
	myStack.printStack()
	myStack.pop()
	myStack.pop()
	myStack.pop()
	myStack.printStack()
	fmt.Println(myStack.read())
}