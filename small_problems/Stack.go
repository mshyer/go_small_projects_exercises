package main
import "fmt"

//Stack
//Push
//Pop
//Array? 
type Stack struct {
	items []int
}

func (s *Stack) push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) pop() int {
	l := len(s.items) - 1
	lastEle := s.items[l]
	s.items = s.items[:l]
	return lastEle
}

func (s *Stack) read() {
	for idx := 0; idx < len(s.items); idx++ {
		fmt.Printf("%d\n", s.items[idx])
	}
}

func main() {
	myStack := Stack{items: []int{1, 2, 3, 4, 5, 6}}
	myStack.push('p')
	myStack.read()
	fmt.Println(myStack.pop())
	myStack.read()
}