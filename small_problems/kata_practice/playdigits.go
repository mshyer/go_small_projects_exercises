// INPUT: 2 integers
// OUTPUT: integer (-1 or the right answer)
// STEPS:
// Split the integer into digits
// 1234
// Get the sum of raising digits to consecutive powers
// if sum % original == 0, return sum / original
// else return -1
package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	root *Node
}

func GetDigitsList(n int) *LinkedList {
	myList := new(LinkedList)
	rootNode := new(Node)
	rootNode.value = n % 10
	n /= 10
	myList.root = rootNode

	for n > 0 {
		newNode := new(Node)
		newNode.value = n % 10
		n /= 10
		newNode.next = myList.root
		myList.root = newNode
	}
	return myList
}

func DigPow(n, p int) int {
	myList := GetDigitsList(n)
	current := myList.root
	sum := 0
	for current != nil {
		sum += int(math.Pow(float64(current.value), float64(p)))
		p++
		current = current.next
	}
	fmt.Println(sum)
	if sum%n == 0 {
		return sum / n
	} else {
		return -1
	}
}

func main() {
	// n := 1204
	// myList := GetDigitsList(n)

	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(92, 1))
}
