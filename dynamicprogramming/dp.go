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

func subSetSum(arr []int, W, n int, t [][]bool) [][]bool {

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

	return t
}

func equalSumPartition(arr []int, sum int, t [][]bool) bool {
	if sum%2 != 0 {
		return false
	} else {
		W := sum / 2
		n := len(arr)
		mat := subSetSum(arr, W, n, t)
		return mat[n][W]
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

// Unbounded Knapsack
// Recursive
func unboundedKnapsackRecursive(wt, val []int, W, n int) int {

	// base condition
	if n == 0 || W == 0 {
		return 0
	}

	// main condition
	if wt[n-1] <= W {
		return int(math.Max(float64(val[n-1]+unboundedKnapsackRecursive(wt, val, W-wt[n-1], n)), float64(unboundedKnapsackRecursive(wt, val, W, n-1))))
	} else {
		return unboundedKnapsackRecursive(wt, val, W, n-1)
	}
}

// DP Code
func unboundedKnapsackDP(wt, val []int, W, n int, t [][]int) int {

	// base condition
	if n == 0 || W == 0 {
		return 0
	}

	if t[n][W] != -1 {
		return t[n][W]
	}

	if wt[n-1] <= W {
		t[n][W] = int(math.Max(float64(val[n-1]+unboundedKnapsackDP(wt, val, W-wt[n-1], n, t)), float64(unboundedKnapsackDP(wt, val, W, n-1, t))))
		return t[n][W]
	} else {
		t[n][W] = unboundedKnapsackDP(wt, val, W, n-1, t)
		return t[n][W]
	}
}

// DP Iterative Code
func unboundedKnapsackIterative(wt, val []int, W, n int, t [][]int) [][]int {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// base condition
			if i == 0 || j == 0 {
				t[i][j] = 0
				continue
			}

			if wt[i-1] <= j {
				t[i][j] = int(math.Max(float64(val[i-1]+t[i][j-wt[i-1]]), float64(t[i-1][j])))
			} else {
				t[i][j] = t[i-1][j]
			}

		}
	}
	return t
}

// Road cutting problem
func roadCuttingDP(wt, val []int, W, n int, t [][]int) [][]int {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// base condition
			if i == 0 || j == 0 {
				t[i][j] = 0
				continue
			}

			if wt[i-1] <= j {
				t[i][j] = int(math.Max(float64(val[i-1]+t[i][j-wt[i-1]]), float64(t[i-1][j])))
			} else {
				t[i][j] = t[i-1][j]
			}

		}
	}
	return t
}

// find max no of ways to choose coins
func maxNoWays(arr []int, W, n int, t [][]int) [][]int {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// base condition
			if i == 0 {
				t[i][j] = 0
				continue
			}

			if j == 0 {
				t[i][j] = 1
				continue
			}

			if arr[i-1] <= j {
				t[i][j] = t[i][j-arr[i-1]] + t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t
}

// minimum no of Coins
func minNoOfCoins(arr []int, W, n int, t [][]int) [][]int {

	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 {
				t[i][j] = 99999
			}
			if j == 0 {
				t[i][j] = 0
			}
		}
	}

	i := 1
	for j := 1; j < W+1; j++ {
		if j%arr[0] == 0 {
			t[i][j] = j / arr[0]
		} else {
			t[i][j] = 99999
		}
	}

	for i := 2; i < n+1; i++ {
		for j := 1; j < W+1; j++ {
			if arr[i-1] <= j {
				t[i][j] = int(math.Min(float64(t[i][j-arr[i-1]]+1), float64(t[i-1][j])))
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t
}

// LONGESt COMMON SUBSEQUENCE Problems
// length of LCS
func LCSRecursive(x, y string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if x[m-1] == y[n-1] {
		return 1 + LCSRecursive(x, y, m-1, n-1)
	} else {
		return int(math.Max(float64(LCSRecursive(x, y, m-1, n)), float64(LCSRecursive(x, y, m, n-1))))
	}
}

// LCS DP Code
func LongestCommonSubsequenceDP(x, y string, m, n int, t [][]int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if t[m][n] != -1 {
		return t[m][n]
	}

	if x[m-1] == y[n-1] {
		t[m][n] = 1 + LongestCommonSubsequenceDP(x, y, m-1, n-1, t)
		return t[m][n]
	} else {
		t[m][n] = int(math.Max(float64(LongestCommonSubsequenceDP(x, y, m-1, n, t)), float64(LongestCommonSubsequenceDP(x, y, m, n-1, t))))
		return t[m][n]
	}
}

// LCS Iterative Code
func LCSIterative(x, y string, m, n int, t [][]int) [][]int {

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {

			// base condition
			if i == 0 || j == 0 {
				t[i][j] = 0
				continue
			}

			if x[i-1] == y[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = int(math.Max(float64(t[i-1][j]), float64(t[i][j-1])))
			}
		}
	}
	return t
}

// longest common substring
func longestCommonSubstring(x, y string, m, n int, t [][]int) int {

	// base condition
	result := 0
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {

			// base condition
			if i == 0 || j == 0 {
				t[i][j] = 0
				continue
			}

			if x[i-1] == y[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
				result = int(math.Max(float64(result), float64(t[i][j])))
			} else {
				t[i][j] = 0
			}
		}
	}
	return result
}

// shortest sub sequence
func shortestSubSequence(x, y string, m, n int) int {

	if m == 0 || n == 0 {
		return 0
	} else {
		scsMat := make([][]int, m+1)
		for i := range scsMat {
			scsMat[i] = make([]int, n+1)
			for j := range scsMat[i] {
				scsMat[i][j] = -1
			}
		}
		scsMat = LCSIterative(x, y, m, n, scsMat)
		l := scsMat[m][n]
		return m + n - l
	}
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
	fmt.Println("Is a subset exist whose sum is equal to given sum", isSubSetSum[n][sum])

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
	arr = []int{1, 2, 7}
	n = len(arr)
	total_sum = 0
	for _, num := range arr {
		total_sum += num
	}
	subSetMat := make([][]bool, n+1)
	for i := range subSetMat {
		subSetMat[i] = make([]bool, total_sum+1)
		for j := range subSetMat[i] {
			subSetMat[i][j] = false
		}
	}
	subSetMat = subSetSum(arr, total_sum, n, subSetMat)
	nums := []int{}
	rangeVal := total_sum/2 + 1
	for j := 0; j < rangeVal; j++ {
		if subSetMat[n][j] {
			nums = append(nums, j)
		}
	}
	min_val := 999999
	for _, num := range nums {
		min_val = int(math.Min(float64(min_val), float64(total_sum-2*num)))
	}
	fmt.Println("The minimum sum difference is", min_val)

	// 2.5 Count number of subsets with given difference
	/*
	 We have to partition the array into two subsets such that if we do the sum of these two subsets and when we find the difference that should be equal to given difference and we have to count these possible subsets.
	*/
	arr = []int{1, 1, 2, 3}
	n = len(arr)
	diff := 1
	total_sum = 0
	for _, num := range arr {
		total_sum += num
	}
	sum = diff + total_sum/2
	countMat = make([][]int, n+1)
	for i := range countMat {
		countMat[i] = make([]int, sum+1)
		for j := range countMat[i] {
			countMat[i][j] = 0
		}
	}
	subsets := countSubSetPartition(arr, sum, n, countMat)
	fmt.Println("No of subsets are", subsets)

	// 3. UNBOUNDED KNAPSACK
	// In unbounded  knapsack if a weight is picked we can pick it again.
	wt = []int{1, 3, 4, 5}
	val = []int{1, 4, 5, 7}
	W = 10
	n = len(arr)
	maxProfit = unboundedKnapsackRecursive(wt, val, W, n)
	fmt.Println("Max profit is", maxProfit)

	// DP code for unbounded knapsack
	unboundMat := make([][]int, n+1)
	for i := range unboundMat {
		unboundMat[i] = make([]int, W+1)
		for j := range unboundMat[i] {
			unboundMat[i][j] = -1
		}
	}
	maxProfit = unboundedKnapsackDP(wt, val, W, n, unboundMat)
	fmt.Println("Max Profit in unbounded knapsack ", maxProfit)

	// Itertive code for knapsack
	unboundMat = make([][]int, n+1)
	for i := range unboundMat {
		unboundMat[i] = make([]int, W+1)
		for j := range unboundMat[i] {
			unboundMat[i][j] = -1
		}
	}
	unboundMat = unboundedKnapsackIterative(wt, val, W, n, unboundMat)
	fmt.Println("Max Profit in unbounded knapsack iterative is", unboundMat[n][W])

	// 3.1 Rod cutting problem
	/*
			We have given a rod of length N and we have to cut the road into length of different pieces and for each piece of length we have a corresponding price array
		    so we will get the profit and we have to maximize the profit. Similar to Unbounded knapsack.
	*/
	n = 8
	length := []int{}
	for i := 1; i < n+1; i++ {
		length = append(length, i)
	}
	prices := []int{1, 5, 9, 13, 14, 17, 17, 20}
	rodMat := make([][]int, n+1)
	for i := range rodMat {
		rodMat[i] = make([]int, n+1)
		for j := range rodMat[i] {
			rodMat[i][j] = 0
		}
	}
	rodMat = roadCuttingDP(length, prices, n, n, rodMat)
	fmt.Println("Rod MAt", rodMat)
	fmt.Println("Max Profit in rod cutting is", rodMat[n][n])

	// 3.2 Coin Change Problem
	/*
		Coin Change Problem
		There are two parts of coin change problem
		i) Max no of ways
		ii) Minimum no of coins required

		Problem Statement: We have given an array of coin and a sum we have to select the coins from the array and if sum up the selected coins it should be equal to given sum so we have
		to find out the maximum number of ways in which we can select the coins and in the second variation we have to find the minimum number of coins required.
	*/

	coin := []int{1, 2, 3}
	sum = 5
	n = len(coin)
	coinMat := make([][]int, n+1)
	for i := range coinMat {
		coinMat[i] = make([]int, sum+1)
		for j := range coinMat[i] {
			coinMat[i][j] = 0
		}
	}
	coinMat = maxNoWays(coin, sum, n, coinMat)
	fmt.Println("Max no of ways are", coinMat[n][sum])

	// min number of coins
	minCoinMat := make([][]int, n+1)
	for i := range minCoinMat {
		minCoinMat[i] = make([]int, sum+1)
		for j := range minCoinMat[i] {
			minCoinMat[i][j] = -1
		}
	}
	minCoinMat = minNoOfCoins(coin, sum, n, minCoinMat)
	fmt.Println("Min No of Coins used are", minCoinMat[n][sum])

	// 4. Longest Common Subsequence
	/*
		We have given two string X and Y of length m and n, we have to find the longest subsequence from these two string and have to print the length of subsequence.
	*/
	x := "abcdgh"
	y := "abedfhr"
	m := len(x)
	n = len(y)
	l := LCSRecursive(x, y, m, n)
	fmt.Println("Length of longest common subsequence", l)

	// LCS DP
	lcsMat := make([][]int, m+1)
	for i := range lcsMat {
		lcsMat[i] = make([]int, n+1)
		for j := range lcsMat[i] {
			lcsMat[i][j] = -1
		}
	}
	l = LongestCommonSubsequenceDP(x, y, m, n, lcsMat)
	fmt.Println("Length of longest common subsequence", l)

	// LCS Iterative
	lcsMatrix := make([][]int, m+1)
	for i := range lcsMatrix {
		lcsMatrix[i] = make([]int, n+1)
		for j := range lcsMatrix[i] {
			lcsMatrix[i][j] = -1
		}
	}
	lcsMatrix = LCSIterative(x, y, m, n, lcsMatrix)
	fmt.Println("Length of longest common subsequence", lcsMatrix[m][n])

	// longest common substring
	a := "abcde"
	b := "abced"
	m = len(a)
	n = len(b)
	lcsSubMat := make([][]int, m+1)
	for i := range lcsSubMat {
		lcsSubMat[i] = make([]int, n+1)
		for j := range lcsSubMat[i] {
			lcsSubMat[i][j] = -1
		}
	}
	l = longestCommonSubstring(a, b, m, n, lcsSubMat)
	fmt.Println("Length of longest common substring", l)

	// shortest common subsequence
	/*
	 We have given two strings and we have to merge these two strings and have to form a new string such that these strings are subsequence of the new string we have so many strings
	 take string of minimum length.
	*/
	a = "geek"
	b = "eke"
	l = shortestSubSequence(a, b, len(a), len(b))
	fmt.Println("Length of scs is", l)

}
