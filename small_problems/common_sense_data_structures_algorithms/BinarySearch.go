package main

import (
	"fmt"
)



type IntTreeNode struct {
	value int;
	leftChild *IntTreeNode;
	rightChild *IntTreeNode;
}

func (i *IntTreeNode) initialize(val int) {
	i.value = val
}

func (i *IntTreeNode) drawTree() {
	fmt.Printf("  %d  \n", i.value)
	fmt.Printf(" /  \\ \n")
	fmt.Printf(" %d  %d \n", i.leftChild.value, i.rightChild.value)
}

func (i* IntTreeNode) search(searchVal int) *IntTreeNode {
	fmt.Println("firing", i)
	if (i.value) == searchVal {
		return i
	} else if searchVal < i.value && i.leftChild != nil {
		return i.leftChild.search(searchVal)
	} else if searchVal > i.value && i.rightChild != nil {
		return i.rightChild.search(searchVal)
	}
	errorTreeNode := new(IntTreeNode)
	errorTreeNode.initialize(-1)
	return errorTreeNode
}

func (i* IntTreeNode) insert(value int) {
	if value == i.value {
		return
	} else if i.leftChild == nil && value < i.value {
		newNode := new(IntTreeNode)
		newNode.initialize(value)
		i.leftChild = newNode
		return
	} else if i.rightChild == nil && value > i.value {
		newNode := new(IntTreeNode)
		newNode.initialize(value)
		i.rightChild = newNode
		return
	} else if value < i.value {
		 i.leftChild.insert(value)
	} else if value > i.value {
		i.rightChild.insert(value)
	}
	return
}

func (nodeToDelete *IntTreeNode) lift(node *IntTreeNode) *IntTreeNode {
	if node.leftChild == nil {
		nodeToDelete.value = node.value
		return node.rightChild
	} else {
		node.leftChild = node.leftChild.lift(nodeToDelete)
		return node
	}
}

func (i *IntTreeNode) delete(value int) *IntTreeNode {
	if value < i.value {
		i.leftChild = i.leftChild.delete(value)
		return i
	} else if value > i.value {
		i.rightChild = i.rightChild.delete(value)
		return i
	}
	if value == i.value {
		if i.leftChild == nil {
			return i.rightChild
		} else if i.rightChild == nil {
			return i.leftChild
		} else {
		  i.rightChild = i.rightChild.lift(i)
		}
	}
	return nil
}

func (node *IntTreeNode) traverseAndPrint() {
	if node.leftChild != nil {
		node.leftChild.traverseAndPrint()
	}
	fmt.Println(node.value)
	if node.rightChild != nil {
		node.rightChild.traverseAndPrint()
	}
}

func (node *IntTreeNode) findGreatest() int {
	if node.rightChild == nil {
		return node.value
	}
	return node.rightChild.findGreatest()
}

func main() {
	node1 := new(IntTreeNode)
	node1.initialize(4)
	node2 := new(IntTreeNode)
	node2.initialize(2)
	node3 := new(IntTreeNode)
	node3.initialize(6)
	node1.leftChild = node2
	node1.rightChild = node3

	// fmt.Println(node1.search(88))
	// fmt.Println(node1.search(6))
	node1.insert(88)
	fmt.Println(node1.findGreatest(), "haha")
	node1.traverseAndPrint()
	node1.delete(88)
	fmt.Println(node1.search(88))
	fmt.Println(node1.findGreatest(), "haha")
}