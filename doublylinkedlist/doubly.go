package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type LinkedList struct {
	head *Node
}

// print list
func (o *LinkedList) printList() {
	if o.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	curr := o.head
	for curr != nil {
		fmt.Print(curr.val, "<---->")
		curr = curr.right
	}
	fmt.Println()
}

// insert at First
func (o *LinkedList) insertAtFirst(data int) {

	node := &Node{
		val:   data,
		left:  nil,
		right: nil,
	}

	if o.head == nil {
		o.head = node
		return
	}

	node.right = o.head
	o.head.left = node
	o.head = node
	return
}

// insert at end
func (o *LinkedList) insertAtEnd(data int) {
	node := &Node{
		val:   data,
		left:  nil,
		right: nil,
	}

	if o.head == nil {
		o.head = node
		return
	}

	curr := o.head
	for curr.right != nil {
		curr = curr.right
	}

	curr.right = node
	node.left = curr
	return

}

// insert before a node
func (o *LinkedList) insertBeforeNode(data, key int) {
	// 0. Check if the linked list is empty
	if o.head == nil {
		fmt.Println("Doubly linked list is empty")
		return
	}

	// 1. Check if we have to insert before a node
	if o.head.val == key {
		o.insertAtFirst(data)
		return
	}

	// 2. traverse the list and find the right position
	var prev *Node
	curr := o.head
	for curr.right != nil && curr.val != key {
		prev = curr
		curr = curr.right
	}

	if curr.right == nil && curr.val != key {
		fmt.Println("Node is present in the list", key)
		return
	}

	// 3. create new node for insertion
	node := &Node{
		val:   data,
		left:  nil,
		right: nil,
	}

	// 3. insert the node at the right position
	prev.right = node
	node.right = curr
	curr.left = node
	node.left = prev
	return

}

// insert after a node
func (o *LinkedList) insertAfterNode(data, key int) {
	// 0. Check if the linked list is empty
	if o.head == nil {
		fmt.Println("Doubly linked list is empty")
		return
	}

	// 1. traverse the list
	curr := o.head
	for curr.right != nil && curr.val != key {
		curr = curr.right
	}

	// 2. either we find the right node or node is not present in the list
	if curr.right == nil {
		if curr.val != key {
			fmt.Println("Node is not present in the list ", key)
			return
		} else {
			o.insertAtEnd(data)
			return
		}
	}

	// 3. create new node for insert
	node := &Node{
		val:   data,
		left:  nil,
		right: nil,
	}

	// 3. insert at the right position
	node.right = curr.right
	curr.right.left = node
	curr.right = node
	node.left = curr
	return

}

func main() {
	dll := &LinkedList{head: nil}
	dll.insertAtFirst(5)
	dll.insertAtFirst(3)
	dll.printList()
	dll.insertAtEnd(9)
	dll.insertAtEnd(11)
	dll.printList()
	dll.insertBeforeNode(7, 9)
	dll.insertBeforeNode(1, 3)
	dll.printList()
	dll.insertAfterNode(13, 11)
	dll.insertAfterNode(8, 7)
	dll.printList()
}
