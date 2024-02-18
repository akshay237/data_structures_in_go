package main

import (
	"fmt"
	"math"
)

func knapsackRecursive(wt, val []int, W, n int) int {
	// Base Condition
	if n == 0 || W == 0 {
		return 0
	}

	// if wt is less than capacity then we have two options
	if wt[n-1] <= W {
		return int(math.Max(float64(val[n-1]+knapsackRecursive(wt, val, W-wt[n-1], n-1)), float64(knapsackRecursive(wt, val, W, n-1))))
	} else {
		return knapsackRecursive(wt, val, W, n-1)
	}
}

func knapsackDP(wt, val []int, W, n int, t [][]int) int {
	// Base Condition
	if n == 0 || W == 0 {
		return 0
	}

	if t[n][W] != -1 {
		return t[n][W]
	}

	if wt[n-1] <= W {
		t[n][W] = int(math.Max(float64(val[n-1]+knapsackDP(wt, val, W-wt[n-1], n-1, t)), float64(knapsackDP(wt, val, W, n-1, t))))
		return t[n][W]
	} else {
		t[n][W] = knapsackDP(wt, val, W, n-1, t)
		return t[n][W]
	}
}

func knapsackIterative(wt, val []int, W, n int, t [][]int) int {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// base condition
			if i == 0 || j == 0 {
				t[i][j] = 0
				continue
			}

			// if wt is less then the capacity then we have two choices
			if wt[i-1] <= j {
				t[i][j] = int(math.Max(float64(val[i-1]+t[i-1][j-wt[i-1]]), float64(t[i-1][j])))
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][W]
}

func subSetSum(arr []int, W, n int, t [][]bool) bool {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// base condition
			if i == 0 {
				t[i][j] = false
			}
			if j == 0 {
				t[i][j] = true
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < W+1; j++ {
			// main condition
			if arr[i-1] <= j {
				t[i][j] = t[i-1][j-arr[i-1]] || t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}

	return t[n][W]
}

func equalSumPartition(arr []int, sum int, t [][]bool) bool {
	if sum%2 != 0 {
		return false
	} else {
		W := sum / 2
		n := len(arr)
		return subSetSum(arr, W, n, t)
	}
}

func countSubSetPartition(arr []int, W, n int, t [][]int) int {

	// Base Initilization condition
	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 {
				t[i][j] = 0
			}
			if j == 0 {
				t[i][j] = 1
			}
		}
	}

	// Main Condition
	for i := 1; i < n+1; i++ {
		for j := 1; j < W+1; j++ {
			if arr[i-1] <= j {
				t[i][j] = t[i-1][j-arr[i-1]] + t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[n][W]
}

func main() {
	// 1. PROBLEMS BASED ON FIXED KNAPSACK

	// 1.1 Recursive approach
	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}
	W := 7
	n := len(val)
	cost := knapsackRecursive(wt, val, W, n)
	fmt.Println("Max profit is", cost)

	// 2. Memoized code for knapsack
	t := make([][]int, n+1)
	for i := range t {
		t[i] = make([]int, W+1)
		for j := range t[i] {
			t[i][j] = -1
		}
	}
	t[n][W] = knapsackDP(wt, val, W, n, t)
	fmt.Println("Max profit is", t[n][W])

	// 3. iterative code for knapsack
	mat := make([][]int, n+1)
	for i := range mat {
		mat[i] = make([]int, W+1)
		for j := range mat[i] {
			mat[i][j] = 0
		}
	}
	maxProfit := knapsackIterative(wt, val, W, n, mat)
	fmt.Println("Max Profit is", maxProfit)

	// 2. SUBSET Problems
	// 2.1 Is there a subset whose sum is equal to given sum
	arr := []int{2, 3, 7, 8, 10}
	n = len(arr)
	sum := 12
	sumMat := make([][]bool, n+1)
	for i := range sumMat {
		sumMat[i] = make([]bool, sum+1)
		for j := range sumMat[i] {
			sumMat[i][j] = false
		}
	}
	isSubSetSum := subSetSum(arr, sum, n, sumMat)
	fmt.Println("Is a subset exist whose sum is equal to given sum", isSubSetSum)

	// 2.2 Equal Sum Partition
	arr = []int{1, 5, 11, 5}
	total_sum := 0
	for _, num := range arr {
		total_sum += num
	}
	isequalSum := equalSumPartition(arr, total_sum, sumMat)
	fmt.Println("Equal Sum Partition subset exists", isequalSum)

	// 2.3 Count subset of array whose sum is equal to given sum
	arr = []int{2, 3, 5, 6, 8, 10}
	sum = 10
	n = len(arr)
	countMat := make([][]int, n+1)
	for i := range countMat {
		countMat[i] = make([]int, sum+1)
		for j := range countMat[i] {
			countMat[i][j] = 0
		}
	}
	noOfSubsets := countSubSetPartition(arr, sum, len(arr), countMat)
	fmt.Println("No of Subsets are whose sum is equal to given sum", noOfSubsets)

	// 2.4 Minimum Subset Sum Difference
	arr = []int{1, 6, 11, 5}
	n = len(arr)
	total_sum = 0
	for _, num := range arr {
		total_sum += num
	}
}
