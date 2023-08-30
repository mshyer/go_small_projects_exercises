package main

import "fmt"

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) prepend (n *node){
	n.next = l.head
	l.head = n
	l.length++
}

func (l linkedList) printListData() {
	current := l.head
	for idx := 0; idx < l.length; idx++ {
		fmt.Println(current.data)
		current = current.next
	}
}

func (l *linkedList) deleteWithValue(value int) {
	toDelete := l.head
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	} 
	for idx := 0; idx < l.length - 1; idx++ {
		if toDelete.next.data == value {
			toDelete.next = toDelete.next.next
			l.length--
			return
		} else {
			toDelete = toDelete.next
		}
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:12}
	node3 := &node{data:66}
	node4 := &node{data:2}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(12)
	myList.deleteWithValue(100)
	myList.printListData()
	myList.deleteWithValue(48)
	myList.printListData()

}