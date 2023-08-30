/*
1: O(N^2)
2: O(N)
3: 
	BEST: O(1)
	AVERAGE:O(N^2/ 2)
	WORST: O(N^2)
4:m Time complexity: O(N)


*/

package main

func containsX(str string) bool {
	for idx := 0; idx < len(str); idx++ {
		if (str[idx] == 'X') {
			return true
		}
	}
	return false
}

func main() {

}