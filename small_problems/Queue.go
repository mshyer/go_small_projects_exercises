package main
import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(n int) {
	q.items = append(q.items, n)
}

func (q *Queue) dequeue() int {
	deq := q.items[0]
	q.items = q.items[1:]
	return deq
}

func (q Queue) read() {
	for idx := 0; idx < len(q.items); idx++ {
		fmt.Printf("%d\n", q.items[idx])
	}
}

func main() {
	myQueue := Queue{items: []int{1, 2, 3, 4, 5}}
	myQueue.enqueue(6)
	myQueue.read()
	myQueue.dequeue()
	myQueue.read()
	// fmt.Printf("%d\n", deq)
	

}