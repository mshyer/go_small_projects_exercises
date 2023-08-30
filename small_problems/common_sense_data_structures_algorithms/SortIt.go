package main

import "fmt"

func InsertionSort(arr *[]int) *[]int{
	for idx, value := range (*arr) {
		if idx == 0 {
			continue
		}
		position := idx - 1
		for position >= 0 {
			if (*arr)[position] > value {
				fmt.Println((*arr)[position], (*arr)[idx])
			  (*arr)[idx] = (*arr)[position]
				position -= 1
			} else {
				fmt.Println(*arr)
				break
			}
			(*arr)[position + 1] = value
		}
	}
	return arr
}

func main() {
	fmt.Println(InsertionSort(&[]int{4, 3, 6, 7, 1 ,2}))

}