package main
import "fmt"

func printit() {
	myNum := 5

	for idx := 0; idx < 10; idx++ {
		if idx == myNum {
			fmt.Println("whee")
			return
		}
		fmt.Println(idx)
	}
}

func main() {
	printit()
}