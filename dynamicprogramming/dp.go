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

func main() {
	// 1. Problems based on fixed Knapsack
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
}
