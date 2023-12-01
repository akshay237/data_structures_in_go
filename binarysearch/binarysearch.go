package main

import "fmt"

func BinarySearch(arr []int, start, end, key int) int {

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func BinarySearchInDesc(arr []int, start, end, key int) int {

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {

	// 1. Binary Search
	arr := []int{1, 3, 5, 6, 8, 10, 12}
	key := 6

	pos := BinarySearch(arr, 0, len(arr)-1, key)
	fmt.Println("Position of key is: ", pos)

	// 2. BinarySearch in Desc Order Sorted Array
	arr = []int{19, 17, 16, 14, 12, 10, 8, 7, 5, 3, 1}
	key = 16

	pos = BinarySearchInDesc(arr, 0, len(arr)-1, key)
	fmt.Println("Position of key iin desc sorted array: ", pos)

}
