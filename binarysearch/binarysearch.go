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

func FirstOccurence(arr []int, start, end, key int) int {
	first := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			first = mid
			end = mid - 1
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return first
}

func LastOccurence(arr []int, start, end, key int) int {
	last := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			last = mid
			start = mid + 1
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return last
}

func NumRotations(arr []int, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		left := (mid + len(arr) - 1) % len(arr) // to avoid out of index
		right := (mid + 1) % len(arr)

		if arr[left] >= arr[mid] && arr[mid] <= arr[right] { // find the min element or max element (for max reverse the conditions)
			return mid
		} else if arr[start] <= arr[mid] {
			start = mid + 1
		} else if arr[mid] <= arr[end] {
			end = mid - 1
		}
	}
	return -1
}

func searchInNearlySorted(arr []int, start, end, key int) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		} else if mid-1 >= start && arr[mid-1] == key {
			return mid - 1
		} else if mid+1 <= end && arr[mid+1] == key {
			return mid + 1
		} else if arr[mid] > key {
			end = mid - 2
		} else {
			start = mid + 1
		}
	}
	return -1
}

func getFloor(arr []int, start, end, key int) int {
	floor := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			end = mid - 1
		} else {
			floor = mid
			start = mid + 1
		}
	}
	return floor
}

func getCeil(arr []int, start, end, key int) int {
	ceil := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			ceil = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ceil
}

func searchInInfiniteArray(arr []int, key int) int {
	start := 0
	end := 1

	for key > arr[end] {
		start = end
		end = end * 2
	}

	return BinarySearch(arr, start, end, key)
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

	// 3. Order Agnostic Binary Search
	arr = []int{1, 5, 7, 9, 11, 13, 15, 17, 19}
	key = 15
	if arr[0] <= arr[len(arr)-1] {
		pos = BinarySearch(arr, 0, len(arr)-1, key)
	} else {
		pos = BinarySearchInDesc(arr, 0, len(arr)-1, key)
	}
	fmt.Println("Position of key in agnostic array is: ", pos)

	// 4. find first occurence
	arr = []int{2, 4, 10, 10, 10, 18, 20}
	key = 10
	first := FirstOccurence(arr, 0, len(arr)-1, key)
	fmt.Println("First occurence is:", first)

	// 5. find last occurence
	arr = []int{2, 4, 10, 10, 10, 18, 20}
	key = 10
	last := LastOccurence(arr, 0, len(arr)-1, key)
	fmt.Println("Last occurence is:", last)

	// 6. Count the no of the occurence of a key in an array
	arr = []int{2, 4, 10, 10, 10, 18, 20}
	key = 10
	count := last - first + 1
	fmt.Println("Occurences of key is:", count)

	// 7. find the number of times an sorted array is rotated
	arr = []int{6, 8, 11, 12, 15, 18, 2, 5}
	times := NumRotations(arr, 0, len(arr)-1)
	fmt.Println("Sorted array is rotated no of times:", times)

	// 8. find an element in a sorted array
	key = 8
	pos1 := BinarySearch(arr, 0, times-1, key)
	pos2 := BinarySearch(arr, times, len(arr)-1, key)
	if pos1 == -1 && pos2 == -1 {
		fmt.Println("Key  is not present in the array")
	} else if pos1 == -1 && pos2 != -1 {
		fmt.Println("Position of element in the array is:", pos2)
	} else if pos1 != -1 && pos2 == -1 {
		fmt.Println("Position of element in the array is:", pos1)
	}

	// 9. Search an element in a nearly sorted array
	arr = []int{2, 4, 10, 12, 15, 18, 20}
	key = 10
	pos = searchInNearlySorted(arr, 0, len(arr)-1, key)
	fmt.Println("Position of key in nearly sorted array is:", pos)

	// 10. Find floor of an element
	arr = []int{1, 2, 3, 4, 8, 10, 12, 15}
	key = 11
	floor := getFloor(arr, 0, len(arr)-1, key)
	fmt.Println("Floor of the key is: ", arr[floor])

	// 11. Find ceil of an element
	arr = []int{1, 2, 3, 4, 8, 10, 12, 15}
	key = 11
	ceil := getCeil(arr, 0, len(arr)-1, key)
	fmt.Println("Ceil of the key is:", arr[ceil])

	// 12. find position of element in an infinite array
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	key = 7
	pos = searchInInfiniteArray(arr, key)
	fmt.Println("Position of element is:", pos)

	// 13. find position of element in an infinite array of 0 and 1
	arr = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	pos = searchInInfiniteArray(arr, 1)
	fmt.Println("Position of element is:", pos)
}
