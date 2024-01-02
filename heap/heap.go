package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	ele, isok := x.(int)
	if !isok {
		return
	}
	*h = append(*h, ele)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 1. Find the kth Largest element in an array
func kthLargest(arr []int, k int) int {
	// 1. create a heap object
	h := &MinHeap{}

	// 2. initalize the heap
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			_ = heap.Pop(h)
		}
	}
	return (*h)[0]
}

// 2. Find the kth smallest element in an array
func kthSmallest(arr []int, k int) int {

	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, -1*arr[i])
		if h.Len() > k {
			_ = heap.Pop(h)
		}
	}
	return -1 * (*h)[0]
}

// 3. Sort a K sorted array
func KSorted(arr []int, k int) []int {
	ans := make([]int, 0)
	h := &MinHeap{}

	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			val := heap.Pop(h)
			ans = append(ans, val.(int))
		}
	}

	for h.Len() > 0 {
		val := heap.Pop(h)
		ans = append(ans, val.(int))
	}
	return ans
}

// 4. k Closet Numbers
type Closet struct {
	distance int
	value    int
}

type ClosetHeap []*Closet

func (h ClosetHeap) Len() int {
	return len(h)
}

func (h ClosetHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h ClosetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ClosetHeap) Push(x interface{}) {
	ele, isok := x.(*Closet)
	if !isok {
		return
	}
	*h = append(*h, ele)
}
func (h *ClosetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KClosetNums(arr []int, x, k int) []int {
	ans := make([]int, 0)
	h := &ClosetHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		dis := int(math.Abs(float64(x - arr[i])))
		cHeap := &Closet{
			distance: -1 * dis,
			value:    arr[i],
		}
		heap.Push(h, cHeap)
		if h.Len() > k {
			_ = heap.Pop(h)
		}
	}

	for h.Len() > 0 {
		val := heap.Pop(h)
		c, isok := val.(*Closet)
		if !isok {
			return ans
		}
		ans = append(ans, c.value)
	}
	return ans
}

// 5. Top k frequent nums
type Freq struct {
	count, val int
}

type FreqHeap []*Freq

func (h FreqHeap) Len() int {
	return len(h)
}

func (h FreqHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h FreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreqHeap) Push(x interface{}) {
	ele, isok := x.(*Freq)
	if !isok {
		return
	}
	*h = append(*h, ele)
}
func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topFreqElements(arr []int, k int) []int {
	countMap := make(map[int]int, 0)
	for _, num := range arr {
		countMap[num]++
	}

	h := &FreqHeap{}
	heap.Init(h)

	fmt.Println("countMap", countMap)
	for key, val := range countMap {
		heap.Push(h, &Freq{val, key})
		if h.Len() > k {
			_ = heap.Pop(h)
		}
	}

	ans := make([]int, 0)
	fmt.Println("top k freq count", h.Len())
	for h.Len() > 0 {
		val := heap.Pop(h)
		v, isok := val.(*Freq)
		if !isok {
			return ans
		}
		ans = append(ans, v.val)
	}
	return ans
}

// 6. Frequency sort array
func freqSort(arr []int) []int {
	countMap := make(map[int]int, 0)
	for _, num := range arr {
		countMap[num]++
	}

	h := &FreqHeap{}
	heap.Init(h)

	for key, val := range countMap {
		heap.Push(h, &Freq{val, key})
	}

	ans := make([]int, h.Len())
	for h.Len() > 0 {
		val := heap.Pop(h)
		ele, isok := val.(*Freq)
		if !isok {
			return ans
		}
		ans[h.Len()] = ele.val
	}

	return ans
}

// 7. Connect ropes to minimize the cost
func connectRopes(arr []int) int {
	cost := 0
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
	}

	for h.Len() >= 2 {
		l1 := heap.Pop(h).(int)
		l2 := heap.Pop(h).(int)
		cost += l1 + l2
		heap.Push(h, l1+l2)
	}
	return cost
}

// 8. Sum between kth Smalles elements
func sumBetweenKthSmallerElements(arr []int, k1, k2 int) int {
	sum := 0
	left := kthSmallest(arr, k1)
	right := kthSmallest(arr, k2)

	for _, val := range arr {
		if val > left && val < right {
			sum += val
		}
	}
	return sum
}

func main() {

	// 1. kth largest element
	arr := []int{7, 10, 4, 3, 20, 15}
	kthlargest := kthLargest(arr, 3)
	fmt.Println("Kth Largest Element is:", kthlargest)

	// 2. kth smallest element
	kthsmallest := kthSmallest(arr, 3)
	fmt.Println("kth Smallest element is:", kthsmallest)

	// 3. Sort a k - sorted array
	arr = []int{6, 5, 3, 2, 8, 10, 9}
	k := 3
	sortedArr := KSorted(arr, k)
	fmt.Println("K Sorted array is:", sortedArr)

	// 4. find k closet number to a given humber
	arr = []int{10, 2, 14, 4, 7, 6}
	x := 5
	k = 3
	ans := KClosetNums(arr, x, k)
	fmt.Println("K closet nums to x are:", ans)

	// 5. top k frequent numbers
	arr = []int{1, 2, 1, 3, 1, 2, 4}
	ans = topFreqElements(arr, 2)
	fmt.Println("Top k Freq elements are:", ans)

	// 6. frequency sort array
	ans = freqSort(arr)
	fmt.Println("Freq sorted array is:", ans)

	// 7. Connect Ropes to Minimize the cost
	arr = []int{1, 2, 3, 4, 5}
	cost := connectRopes(arr)
	fmt.Println("Minimum cost to connect ropes is:", cost)

	// 8. sum b/w kth smallest element
	arr = []int{1, 3, 12, 5, 15, 11}
	k1 := 3
	k2 := 6
	sum := sumBetweenKthSmallerElements(arr, k1, k2)
	fmt.Println("Sum b/w kth smallest element is:", sum)
}
