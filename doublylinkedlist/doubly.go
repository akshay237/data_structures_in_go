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

// delete from first
func (o *LinkedList) deleteAtFirst() {

	if o.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	o.head = o.head.right
	o.head.left = nil
	fmt.Println("Deleted Node is: ", o.head.val)
	return
}

// delete from end
func (o *LinkedList) deleteAtEnd() {

	if o.head == nil {
		fmt.Println("linked list is empty")
		return
	}

	var prev *Node
	curr := o.head
	for curr.right != nil {
		prev = curr
		curr = curr.right
	}

	prev.right = nil
	fmt.Println("Deleted Node is: ", curr.val)
	return
}

// delete specific node in list
func (o *LinkedList) deleteGivenNode(data int) {

	// 1. Check if the list is empty
	if o.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	// 2. if the node to be deleted is head node
	if o.head.val == data {
		o.deleteAtFirst()
		return
	}

	// 3. traverse the list to find the node
	var prev *Node
	curr := o.head
	for curr.right != nil && curr.val != data {
		prev = curr
		curr = curr.right
	}

	// 4. if the node to be deleted is last node
	if curr.right == nil {
		if curr.val == data {
			o.deleteAtEnd()
			return
		} else {
			fmt.Println("Node is not present in the list")
			return
		}
	}

	// 5. delete the node from other positions
	prev.right = curr.right
	curr.right.left = prev
	fmt.Println("Deleted node is: ", curr.val)
	return
}

// search the node in the list
func (o *LinkedList) searchNode(key int) bool {
	if o.head == nil {
		fmt.Println("Link list is empty, so node is not present")
		return false
	}

	curr := o.head
	for curr.right != nil && curr.val != key {
		curr = curr.right
	}

	if curr.val == key {
		return true
	}

	return false
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
	dll.deleteAtFirst()
	dll.printList()
	dll.deleteAtEnd()
	dll.printList()
	dll.deleteGivenNode(11)
	dll.printList()
	isok := dll.searchNode(9)
	fmt.Println("Node exist: ", isok)
}
