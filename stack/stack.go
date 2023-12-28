package main

import (
	"fmt"
	"log"
	"slices"
)

type Stack struct {
	st []int
	n  int
}

func (s *Stack) isFull() bool {
	if len(s.st) == s.n {
		return true
	}
	return false
}

func (s *Stack) isEmpty() bool {
	if len(s.st) == 0 {
		return true
	}
	return false
}

func (s *Stack) push(ele int) []int {
	if s.isFull() {
		log.Println("Stack is full can't add more element to it")
		return s.st
	}
	s.st = append(s.st, ele)
	return s.st
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		log.Println("Stack is empty can not pop the element")
		return -1
	}
	idx := len(s.st) - 1
	ele := s.st[idx]
	s.st = slices.Delete(s.st, idx, idx+1)
	return ele
}

func (s *Stack) getTop() int {
	if s.isEmpty() {
		log.Println("Stack is empty can't get the top")
		return -1
	} else if s.isFull() {
		return s.st[s.n-1]
	} else {
		return s.st[len(s.st)-1]
	}
}

func NewStack(n int) *Stack {
	return &Stack{
		st: make([]int, 0),
		n:  n,
	}
}

func stackImplementation() {
	n := 10
	s := NewStack(n)
	stack := s.push(5)
	fmt.Println("Stack is:", stack)
	stack = s.push(7)
	stack = s.push(9)
	fmt.Println("Stack is:", stack)
	top := s.getTop()
	fmt.Println("Top of stack is:", top)
	popele := s.pop()
	fmt.Println("Popped ele:", popele)
	fmt.Println("Stack after element popped is:", s.st)
}

// Stack based Problems
// 1. Next Greater element from right (NGR) or Next Largest Element()
func NGR(arr []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < arr[i] {
			idx := len(stack) - 1
			stack = slices.Delete(stack, idx, idx+1)
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ele := stack[len(stack)-1]
			ans[i] = ele
		}
		stack = append(stack, arr[i])
	}
	return ans
}

func main() {

	//1. find NGR of an array
	arr := []int{1, 3, 2, 4}
	ngr := NGR(arr)
	fmt.Println("NGR of arr is:", ngr)
}
