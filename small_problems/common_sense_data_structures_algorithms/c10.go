package main

/* 
1: if low > high
2: Stack overflow for even values of n. For odd values, youd get lowball results
3: 


*/
import (
	"fmt"
)




func sum(low, high int) int {
	if high == low {
		return low
	}
	return high + sum(low, high - 1)
}

func printAllNumbers(nestArr []interface{}) {
	for _, item := range nestArr {
		switch v := item.(type) {
		default:
			fmt.Printf("val is of an unknown type %T\n", v)
		case int:
			fmt.Printf("%d\n", item)
		case []int:
			for idx := 0; idx < len(v); idx++ {
				fmt.Printf("%d\n", v[idx])
			}
		case []interface{}:
			if subArr, ok := item.([]interface{}); ok {
				printAllNumbers(subArr)
			}
		}
	}
}


func main() {
	fmt.Println("%v\n", sum(1, 10))

	myArr := []interface{}{
		1, 2, 3,
		[]int{4, 5, 6},
		7,
		[]interface{}{
			8,
			[]interface{}{
				9, 10, 11,
				[]int{12, 13, 14},
			},
		},
		[]interface{}{
			15, 16, 16, 18, 19,
			[]interface{}{
				20, 21, 22,
				[]interface{}{
					23, 24, 25,
					[]int{26, 27, 29},
					30, 31,
				},
				32,
			},
			33,
		},
	}
	printAllNumbers(myArr)
}

