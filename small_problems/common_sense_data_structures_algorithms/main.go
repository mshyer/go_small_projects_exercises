package main

import (
	"fmt"
)

var pr = fmt.Println

func main() {
	fmt.Println("myArr: ", []int{2, 4, 6, 8, 10, 12, 13})
	counter := 0
	n := 100000
	for n > 1 {
		counter += 1
		n /= 2
	}
	pr(counter)
	myStr := "hello"
	myStr = myStr[:2] + "x" + myStr[2:]
	pr(myStr)
}