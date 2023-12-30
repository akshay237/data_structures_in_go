package main

import (
	"fmt"
	"log"
	"math"
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

// 2. find Next Greater Element from left
func NGL(arr []int) []int {
	stack := make([]int, 0)
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < arr[i] {
			idx := len(stack) - 1
			stack = slices.Delete(stack, idx, idx+1)
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1])
		}
		stack = append(stack, arr[i])
	}
	return ans
}

// 3.find next smallest element from right
func NSR(arr []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] > arr[i] {
			idx := len(stack) - 1
			stack = slices.Delete(stack, idx, idx+1)
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return ans
}

// 4. find next smallest element from the left
func NSL(arr []int) []int {
	stack := make([]int, 0)
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > arr[i] {
			idx := len(stack) - 1
			stack = slices.Delete(stack, idx, idx+1)
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1])
		}
		stack = append(stack, arr[i])
	}
	return ans
}

// 5. Stock Span Problem where we have given price of a stock for each day we have to diff consecutive no of days the price of stock is less than today
func stockSpan(arr []int) []int {
	type tuple struct {
		idx   int
		value int
	}
	stack := make([]tuple, 0)
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1].value < arr[i] {
			index := len(stack) - 1
			stack = slices.Delete(stack, index, index+1)
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1].idx)
		}
		t := tuple{
			idx:   i,
			value: arr[i],
		}
		stack = append(stack, t)
	}
	return ans
}

// 6. Max area of histogram
/*
We have given an array height where each index holds height of a building and if we form building with this height and width 1 we will get a histogram and we have to print the
maximum area out of it.
heights = [6, 2, 5, 4, 5, 1, 6]
out = 12
We have to find the width by finding the NSR and NSL of indexs
*/

func NearestSmallestRight(arr []int) []int {
	type tuple struct {
		idx   int
		value int
	}
	stack := make([]tuple, 0)
	ans := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1].value >= arr[i] {
			index := len(stack) - 1
			stack = slices.Delete(stack, index, index+1)
		}
		if len(stack) == 0 {
			ans[i] = len(arr)
		} else {
			ans[i] = stack[len(stack)-1].idx
		}
		t := tuple{
			idx:   i,
			value: arr[i],
		}
		stack = append(stack, t)
	}
	return ans
}

func NearestSmallestLeft(arr []int) []int {
	type tuple struct {
		idx   int
		value int
	}
	stack := make([]tuple, 0)
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1].value >= arr[i] {
			index := len(stack) - 1
			stack = slices.Delete(stack, index, index+1)
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1].idx)
		}
		t := tuple{
			idx:   i,
			value: arr[i],
		}
		stack = append(stack, t)
	}
	return ans
}

func MAH(height []int) int {
	// 1. stores indexes of smallest element from left
	left := NearestSmallestLeft(height)
	//fmt.Println("Left is:", left)
	// 2. stores indexes of smallest element from right
	right := NearestSmallestRight(height)
	// fmt.Println("Right is:", right)
	// 3. find the width
	width := make([]int, len(height))
	for idx, _ := range height {
		width[idx] = right[idx] - left[idx] - 1
	}
	// 4. find the area
	var max_area int
	for idx, _ := range height {
		max_area = int(math.Max(float64(max_area), float64(height[idx]*width[idx])))
	}
	return max_area
}

// 7. Find max area in a binary matrix
func maxAreaInMatrix(mat [][]int) int {
	max_area := 0
	height := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] = height[j] + mat[i][j]
			}
		}
		fmt.Println("Height is:", height)
		area := MAH(height)
		fmt.Println("Area of each histogram is:", area)
		max_area = int(math.Max(float64(max_area), float64(area)))
	}
	return max_area
}

// 8. Rain Water Tapping Problem
func rainTapWater(arr []int) int {
	maxL := make([]int, len(arr))
	maxR := make([]int, len(arr))

	// assign first and last value from arr
	maxL[0] = arr[0]
	maxR[len(arr)-1] = arr[len(arr)-1]

	for i := 1; i < len(arr); i++ {
		maxL[i] = int(math.Max(float64(maxL[i-1]), float64(arr[i])))
	}

	for i := len(arr) - 2; i >= 0; i-- {
		maxR[i] = int(math.Max(float64(maxR[i+1]), float64(arr[i])))
	}

	waterUnits := 0
	for i := 0; i < len(arr); i++ {
		unit := int(math.Min(float64(maxL[i]), float64(maxR[i]))) - arr[i]
		waterUnits += unit
	}
	return waterUnits
}

// 9. Min Stack with O(1) space
type MinStack struct {
	st     []int
	minEle int
}

func (m *MinStack) push(ele int) []int {
	if len(m.st) == 0 {
		m.minEle = ele
		m.st = append(m.st, ele)
	} else {
		if ele > m.minEle {
			m.st = append(m.st, ele)
		} else {
			m.st = append(m.st, 2*ele-m.minEle)
			m.minEle = ele
		}
	}
	return m.st
}

func (m *MinStack) pop() int {
	if len(m.st) == 0 {
		fmt.Println("min stack is empty")
		return -1
	} else {
		idx := len(m.st) - 1
		val := m.st[idx]
		m.st = slices.Delete(m.st, idx, idx+1)
		if val < m.minEle {
			e := m.minEle
			m.minEle = 2*m.minEle - val
			return e
		}
		return val
	}
}

func NewMinStack() *MinStack {
	return &MinStack{
		st:     make([]int, 0),
		minEle: 0,
	}
}

func main() {

	// 1. find NGR of an array
	arr := []int{1, 3, 2, 4}
	ngr := NGR(arr)
	fmt.Println("NGR of arr is:", ngr)

	// 2. find ngl of an array
	ngl := NGL(arr)
	fmt.Println("NGL of arr is:", ngl)

	// 3. find the NSR of an array
	nsr := NSR(arr)
	fmt.Println("NSR of arr is:", nsr)

	// 4. find the NSL of an array
	nsl := NSL(arr)
	fmt.Println("NSL of arr is:", nsl)

	// 5. stock span problem
	arr = []int{100, 80, 60, 70, 60, 75, 85}
	span := stockSpan(arr)
	for i := 0; i < len(span); i++ {
		span[i] = i - span[i]
	}
	fmt.Println("Stock span days are:", span)

	// 6. find maximum area of histogram
	height := []int{6, 2, 5, 4, 5, 1, 6}
	max_area := MAH(height)
	fmt.Println("Maximum area of histogram is:", max_area)

	// 7. max area in a binary matrix
	matrix := [][]int{{0, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 0, 0}}
	max_area = maxAreaInMatrix(matrix)
	fmt.Println("Max area in a binary matrix is: ", max_area)

	// 8. rain water tapping
	height = []int{3, 0, 0, 2, 0, 4}
	waterUnits := rainTapWater(height)
	fmt.Println("Water unit tapped are:", waterUnits)

	// 9. implement stack with o(1)
	m := NewMinStack()
	ms := m.push(9)
	ms = m.push(7)
	ms = m.push(10)
	fmt.Println("Min Stack: ", ms)
	fmt.Println("Min Ele: ", m.minEle)
	p := m.pop()
	fmt.Println("Popped Ele:", p)
	fmt.Println("Min Stack after ele popped:", m.st)
	fmt.Println("Min Ele is:", m.minEle)
	ms = m.push(3)
	fmt.Println("Min Stack: ", ms)
	fmt.Println("Min Ele: ", m.minEle)
}
