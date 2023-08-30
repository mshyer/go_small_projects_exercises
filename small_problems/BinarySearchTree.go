package main
import "fmt"

type TreeNode struct{
	data int
	lChild *TreeNode
	rChild *TreeNode
	parent *TreeNode
}

// type BinarySearchTree struct {
// 	root *TreeNode
// 	length int
// }

// func (bst *BinarySearchTree) init(root *TreeNode) {
// 	bst.length = 1
// 	bst.root = root
// }

func (root *TreeNode) insert(leaf *TreeNode) {
	//Using recursion
	//Input: TreeNode pointer
	//Output: mutate Binary Tree
	//Data Structure: Recursion
	//Problem: Stop condition? if no child where you otherwise recurse
	// Keep recursing until you hit that stop condition
	// if leaf < currentNode, recurse on the left child
	// if leaf > currentNode, recurse on the right child
	currentNode := root
	if leaf.data < currentNode.data {
		if currentNode.lChild == nil {
			currentNode.lChild = leaf
		} else {
			currentNode.lChild.insert(leaf)
		}
	} else if leaf.data > currentNode.data {
		if currentNode.rChild == nil {
			currentNode.rChild = leaf
		} else {
			currentNode.rChild.insert(leaf)
		}
	}
}


	//Using loop structure
	//While the current node has children
	// if the current node is equal to the leavNode's value, return
	// if the leaf node value is greater than current
	//    if the current has right child, set current to rightCHild
  //      else set current's right child to leaf, then return
	// if current is less than leaf, 
	      // if current has left child set current to left child
				// else set current's left child to leaf, then return
	//OUT OF WHILE LOOP
	//  if the leaf is aqual to currentNode, return
	//  if leaf less than current, set current left to leaf
	//  if leaf greater than current. set current greater to leaf
// 	currentNode := bst.root
// 	for currentNode.lChild != nil || currentNode.rChild != nil {

// 		if leaf.data == currentNode.data {
// 			return
// 		} else if leaf.data < currentNode.data {
// 			if currentNode.lChild == nil {
// 				currentNode.lChild = leaf
// 				bst.length += 1
// 				return
// 			} else {
// 				currentNode = currentNode.lChild
// 			}
// 		} else {
// 			if currentNode.rChild == nil {
// 				currentNode.rChild = leaf
// 				bst.length += 1
// 				return
// 			} else {
// 				currentNode = currentNode.rChild
// 			}
// 		}
// 	}
// 	if leaf.data == currentNode.data {
// 		fmt.Println("retirning")
// 		return
// 	} else if leaf.data < currentNode.data {
// 		currentNode.lChild = leaf
// 		bst.length += 1
// 	} else if leaf.data > currentNode.data {
// 		currentNode.rChild = leaf
// 		bst.length += 1
// 	} else {
// 		fmt.Println("Some error occurred.")
// 	}
// }

func (currentNode TreeNode) search(val int) bool {
	//input: int
	//Output: boolean
	//Data structure: recursion
	//Target Complexity: O(Log N)
	//Stop condition: if data == search val return true
	// else if less, recurse on left, unless if none to the left, return false
	// if greater, recurse on rights, unless none to right, return false

	if val > currentNode.data {
		if currentNode.rChild == nil {
			return false
		}
		return currentNode.rChild.search(val)
	} else if val < currentNode.data {
		if currentNode.lChild == nil {
			return false
		}
		return currentNode.lChild.search(val)
	}
	return true
}

func (bst TreeNode) printTree(current *TreeNode) {
	if current.lChild != nil {
		bst.printTree(current.lChild)
	}
	fmt.Printf("%d\n", current.data)

	if current.rChild != nil {
		bst.printTree(current.rChild)
	}

}

func main() {
	node2 := TreeNode{data: 25}
	node3 := TreeNode{data: 75}
	node4 := TreeNode{data: 22}
	node5 := TreeNode{data: 88}
	node6 := TreeNode{data: 12}
	node7 := TreeNode{data: 12}
	node8 := TreeNode{data: 66}
	node9 := TreeNode{data: 99}
	node10 := TreeNode{data: 17}
	node11 := TreeNode{data: 21}
	node12 := TreeNode{data: 82}
	node13 := TreeNode{data: 49}
	node14 := TreeNode{data: 46}
	node15 := TreeNode{data: 4}

	// treeRoot := TreeNode{data: 50}
	myTree := TreeNode{data: 25}
	// myTree.init(&treeRoot)
	myTree.insert(&node2)
	myTree.insert(&node3)
	myTree.insert(&node4)
	myTree.insert(&node5)
	myTree.insert(&node6)
	myTree.insert(&node7)
	myTree.insert(&node8)
	myTree.insert(&node9)
	myTree.insert(&node10)
	myTree.insert(&node11)
	myTree.insert(&node12)
	myTree.insert(&node13)
	myTree.insert(&node14)
	myTree.insert(&node15)

	// node14 := TreeNode{data: 22}
	// node15 := TreeNode{data: 22}

	// myTree.insert(&node14)
	// myTree.insert(&node15)


	myTree.printTree(&myTree)
	fmt.Printf("%v\n", myTree.search(22))
	fmt.Printf("%v\n", myTree.search(99))
	fmt.Printf("%v\n", myTree.search(333))
}