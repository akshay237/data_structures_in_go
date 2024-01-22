package main

import (
	"fmt"
	"math"
)

type BST struct {
	key   *int
	left  *BST
	right *BST
}

// height of a binary tree
func height(node *BST) int {
	if node == nil {
		return 0
	}
	lh := height(node.left)
	rh := height(node.right)
	return 1 + int(math.Max(float64(lh), float64(rh)))
}

// insert a node in the binary search tree
func (o *BST) insert(data int) {
	// if key is nil insert at the root
	if o.key == nil {
		o.key = &data
		return
	}

	// if key is greater than data than insert the node in left subtree otherwise in the right subtree
	if *o.key >= data {
		if o.left != nil {
			o.left.insert(data)
		} else {
			o.left = &BST{
				key:   &data,
				left:  nil,
				right: nil,
			}
		}
	} else {
		if o.right != nil {
			o.right.insert(data)
		} else {
			o.right = &BST{
				key:   &data,
				left:  nil,
				right: nil,
			}
		}
	}
}

// search a node in the binary search tree
func (o *BST) searchNode(data int) {

	if o.key == nil {
		fmt.Println("Binary Search Tree is empty")
		return
	}

	if *o.key == data {
		fmt.Println("Node found in the tree")
		return
	}

	if *o.key > data {
		if o.left != nil {
			o.left.searchNode(data)
		} else {
			fmt.Println("Node is not present in the tree")
			return
		}
	} else {
		if o.right != nil {
			o.right.searchNode(data)
		} else {
			fmt.Println("Node is not present in the tree")
			return
		}
	}
}

// get the min node from binary search tree
func (o *BST) getMinNode() *BST {
	curr := o.left
	for curr != nil {
		curr = curr.left
	}
	return curr
}

// get the max node from binary search tree
func (o *BST) getMaxNode() *BST {
	curr := o.right
	for curr != nil {
		curr = curr.right
	}
	return curr
}

// delete node from a tree
func (o *BST) deleteNode(data int) *BST {

	// 1. check if the tree is empty or not
	if o.key == nil {
		fmt.Println("Binary Search Tree is empty")
		return o
	}

	// 2. Check if the data is greater than or less than root node
	if *o.key > data {
		if o.left != nil {
			o.left = o.left.deleteNode(data)
		} else {
			fmt.Println("Node is not present in the tree")
			return o
		}
	} else if *o.key < data {
		if o.right != nil {
			o.right = o.right.deleteNode(data)
		} else {
			fmt.Println("Node is not present in the tree")
			return o
		}
	} else {
		// 3.1. node has no childs
		if o.left == nil && o.right == nil {
			return &BST{}
		}

		// 3.2. node which is deleted have only right child
		if o.left == nil {
			temp := o.right
			o = nil
			return temp
		}

		// 3.3. node has only left child
		if o.right == nil {
			temp := o.left
			o = nil
			return temp
		}

		// 3.4. node has two child find the min node from right sub tree and make it root node
		temp := o.right.getMinNode()
		o.key = temp.key
		o.right = o.right.deleteNode(*temp.key)
		// return o
	}

	return o
}

// preorder traversal of the binary search tree
func (o *BST) preOrder() {
	if o.key != nil {
		fmt.Print(*o.key, "-->")
		if o.left != nil {
			o.left.preOrder()
		}
		if o.right != nil {
			o.right.preOrder()
		}
	}
}

// in order traversal of binary search tree
func (o *BST) inOrder() {
	if o.key != nil {
		if o.left != nil {
			o.left.inOrder()
		}
		fmt.Print(*o.key, "-->")
		if o.right != nil {
			o.right.inOrder()
		}
	}
}

// post order traversal of binary search tree
func (o *BST) postOrder() {
	if o.key != nil {
		if o.left != nil {
			o.left.postOrder()
		}
		if o.right != nil {
			o.right.postOrder()
		}
		fmt.Print(*o.key, "-->")
	}
}

func main() {
	val := 5
	root := &BST{
		key:   &val,
		left:  nil,
		right: nil,
	}
	root.insert(3)
	root.insert(4)
	root.insert(2)
	root.insert(7)
	root.insert(6)
	root.insert(8)
	fmt.Println("PreOrder traversal is: ")
	root.preOrder()
	fmt.Println()
	fmt.Println("InOrder traversal is: ")
	root.inOrder()
	fmt.Println()
	fmt.Println("PostOrder traversal is: ")
	root.postOrder()
	fmt.Println()
	root.searchNode(7)
	root.deleteNode(2)
	fmt.Println("PreOrder traversal after deleting node 2 is: ")
	root.preOrder()

}
