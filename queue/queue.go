package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.isEmpty() {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
func main() {
	// Create a new queue
	myQueue := Queue{}

	// Enqueue elements
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	// Dequeue and print elements
	fmt.Println("Dequeued:", myQueue.Dequeue())
	fmt.Println("Dequeued:", myQueue.Dequeue())

	// Enqueue more elements
	myQueue.Enqueue(4)
	myQueue.Enqueue(5)

	// Print size of the queue
	fmt.Println("Queue size:", myQueue.Size())

	// Check if the queue is empty
	fmt.Println("Is empty?", myQueue.isEmpty())
}
