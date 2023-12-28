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

func peakElement(arr []int, start, end int) int { // peak element greater than from both left and right side element
	for start <= end {
		mid := start + (end-start)/2
		if mid > 0 && mid < len(arr)-1 {
			if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
				return mid
			} else if arr[mid] > arr[mid-1] {
				start = mid + 1
			} else if arr[mid] > arr[mid+1] {
				end = mid - 1
			}
		} else if mid == 0 {
			if arr[0] > arr[1] {
				return 0
			} else {
				return 1
			}
		} else if mid == len(arr)-1 {
			if arr[len(arr)-1] > arr[len(arr)-2] {
				return len(arr) - 1
			} else {
				return len(arr) - 2
			}
		}
	}
	return -1
}

func searchInMatrix(mat [][]int, rows, cols, key int) (int, int) {
	i := 0
	j := cols - 1
	for i >= 0 && i < rows && j >= 0 && j < cols {
		if mat[i][j] == key {
			return i, j
		} else if mat[i][j] > key {
			j = j - 1
		} else if mat[i][j] < key {
			i = i + 1
		}
	}
	return -1, -1
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

	// 14. find peak element in the array or max in Biotonic Array (biotonic - strictly increasing and strictly decreasing)
	arr = []int{10, 20, 30, 40, 48, 57, 66, 9, 5, 3, 2, 1}
	peakPos := peakElement(arr, 0, len(arr)-1)
	fmt.Println("Peak element position is:", peakPos)

	// 15. search in biotonic array, after getting peak element apply BS and BSinDesc to get the position
	arr = []int{10, 20, 30, 40, 48, 57, 66, 9, 5, 3, 2, 1}
	key = 3
	pos1 = BinarySearch(arr, 0, peakPos, key)
	pos2 = BinarySearchInDesc(arr, peakPos+1, len(arr)-1, key)
	if pos1 == -1 && pos2 == -1 {
		fmt.Println("Key not exist in the biotonic array")
	} else if pos1 != -1 && pos2 == -1 {
		fmt.Println("Key in biotonic array is present at the position", pos1)
	} else if pos1 == -1 && pos2 != -1 {
		fmt.Println("Key in biotonic array is present at the position", pos2)
	}

	// 16. search in a row wise and column wise sorted matrix
	mat := [][]int{{10, 20, 30, 40}, {15, 25, 35, 45}, {27, 29, 37, 48}, {32, 33, 39, 50}}
	key = 29
	rows := len(mat) - 1
	cols := len(mat[0])
	i, j := searchInMatrix(mat, rows, cols, key)
	fmt.Println("Key present at index in mat[", i, "][", j, "]")
}
