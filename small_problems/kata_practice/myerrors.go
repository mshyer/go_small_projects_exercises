package main

import (
	// "errors"
	"fmt"
)

func divide(n, x int) (int, error) {
	if x == 0 {
		// panic("Invalid divisor")
		panic("divided by zero")
	}
	return n / x, nil
}

func add(a, b interface{}) int {
	return a.(int) + b.(int)
}

type Processor func(s string) string

func processString(processor Processor, s string) string {
	return processor(s)
}

func concatenate(s string) string {
	return s + "weeee"
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recovered from:", err)
	// 	}
	// }()

	fmt.Println(process(concatenate, "yay"))

}
