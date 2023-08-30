/*
1:
N Elements  O(N)    o(Log N)    O(N^2)
100         100     7           10000
2000        2000    12          4 million

2: 16 

3: O(N^2)




*/
package main

import (
  "fmt"
)

func greatestNumber(myNums []int) (greatestNumber int) {
	for idx := 0; idx < len(myNums); idx++ {
		if myNums[idx] > greatestNumber {
			greatestNumber = myNums[idx]
		}
	}
	return
}

func main() {

}
