package main

import (
	"fmt"
	"math"
	"slices"
)

//1. Maximum Sum sub array of size k.
//Problem: We have given an array of length n and we have given a vriable k.

func maxSumSubArray(arr []int, k int) (int, int, int) {
	i := 0
	j := 0
	ans := 0
	total := 0
	start := 0
	end := 0
	for j < len(arr) {
		total += arr[j]
		if j-i+1 < k {
			j += 1
		} else if j-i+1 == k {
			ans = int(math.Max(float64(ans), float64(total)))
			if ans == total {
				start = i
				end = j
			}
			total -= arr[i]
			i += 1
			j += 1
		}
	}
	return ans, start, end
}

// 2. First negative integer in every window of size k
func firstNegativeInt(arr []int, k int) []int {
	ans := make([]int, 0)
	temp := make([]int, 0)
	i := 0
	j := 0
	for j < len(arr) {
		if arr[j] < 0 {
			temp = append(temp, arr[j])
		}
		if j-i+1 < k {
			j += 1
		} else if j-i+1 == k {
			if len(temp) == 0 {
				ans = append(ans, 0)
			} else {
				ans = append(ans, temp[0])
				if arr[i] == temp[0] {
					fmt.Println("Temp", temp)
					temp = temp[1:]
				}
			}
			fmt.Println("Ans", ans)
			i += 1
			j += 1
		}
	}
	return ans
}

// 3. Count the occurences of anagrams in a string
func countOccurences(countMap map[string]int, s, p string) int {
	ans := 0
	count := len(countMap)
	i := 0
	j := 0
	k := len(p)
	for j < len(s) {
		ch := string(s[j])
		if val, isok := countMap[ch]; isok {
			val -= 1
			countMap[ch] = val
			if val == 0 {
				count--
			}
		}
		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			if count == 0 {
				ans++
			}
			c := string(s[i])
			if val, isok := countMap[c]; isok {
				val += 1
				countMap[c] = val
				if val == 1 {
					count++
				}
			}
			i++
			j++
		}
	}
	return ans
}

// 4. find maximum of all sub array of size k
func maxOfSizeK(arr []int, k int) []int {
	temp := make([]int, 0)
	ans := make([]int, 0)
	i := 0
	j := 0
	for j < len(arr) {
		for len(temp) > 0 && temp[len(temp)-1] < arr[j] {
			idx := len(temp) - 1
			temp = slices.Delete(temp, idx, idx+1)
		}
		temp = append(temp, arr[j])

		if j-i+1 < k {
			j += 1
		} else if j-i+1 == k {
			ans = append(ans, temp[0])
			if arr[i] == temp[0] {
				temp = slices.Delete(temp, 0, 1)
			}
			i += 1
			j += 1
		}
	}
	return ans
}

// Variable Size Window
// 1. Largest sub array of sum k
func largestSubArrayOfSumK(arr []int, k int) (int, int) {
	start := 0
	end := 0
	i := 0
	j := 0
	total := 0
	ans := 0
	for j < len(arr) {
		total += arr[j]
		if total < k {
			j += 1
		} else if total == k {
			ans = int(math.Max(float64(ans), float64(j-i+1)))
			if ans == j-i+1 {
				start, end = i, j
			}
			j += 1
		} else {
			for total > k {
				total -= arr[i]
				i += 1
			}
			j += 1
		}
	}
	return start, end
}

// 2. Longest sub string with k unique characters
func longSubStrWithKUnique(s string, k int) (int, int) {
	start := 0
	end := 0
	i := 0
	j := 0
	ans := 0
	countMap := make(map[string]int)

	for j < len(s) {
		// add the calculation for each character in the string
		val, isok := countMap[string(s[j])]
		if isok {
			val += 1
		} else {
			val = 1
		}
		countMap[string(s[j])] = val

		if len(countMap) < k {
			j += 1
		} else if len(countMap) == k {
			ans = int(math.Max(float64(ans), float64(j-i+1)))
			if ans == j-i+1 {
				start, end = i, j
			}
			j += 1
		} else {
			// remove the calculation of the character once the window size increased
			for len(countMap) > k {
				val, isok := countMap[string(s[i])]
				if isok {
					val = val - 1
					if val == 0 {
						delete(countMap, string(s[i]))
					} else {
						countMap[string(s[i])] = val
					}
				}
				i += 1
			}
			j += 1
		}
	}
	fmt.Println("Length of longest sub str is:", ans)
	return start, end
}

// 3. longest sub string with all unique
func longSubStringWithAllUnique(s string) (int, int) {
	start := 0
	end := 0
	i := 0
	j := 0
	ans := 0
	countMap := make(map[string]int)

	for j < len(s) {
		val, isok := countMap[string(s[j])]
		if isok {
			val += 1
		} else {
			val = 1
		}
		countMap[string(s[j])] = val

		if len(countMap) == j-i+1 {
			ans = int(math.Max(float64(ans), float64(j-i+1)))
			if ans == j-i+1 {
				start, end = i, j
			}
			j += 1
		} else if len(countMap) < j-i+1 {
			for len(countMap) < j-i+1 {
				val, isok := countMap[string(s[i])]
				if isok {
					val -= 1
					if val == 0 {
						delete(countMap, string(s[i]))
					} else {
						countMap[string(s[i])] = val
					}
				}
				i += 1
			}
			j += 1
		}
	}
	return start, end
}

// 4. find a substring of maximum length i.e a palindrome from a string
func isPalindrome(s string) bool {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return s == str
}

func palindromeSubStrOfMaxLen(s string) (int, int) {
	ans := 0
	start := 0
	end := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			isPali := isPalindrome(s[i:j])
			if isPali {
				ans = int(math.Max(float64(ans), float64(j-i+1)))
				if ans == j-i+1 {
					start, end = i, j
				}
			}
		}
	}
	return start, end
}

func main() {

	// 1. maximum sum in a sub array of size k
	arr := []int{2, 6, 7, 9, 1, 5, 23}
	k := 3
	sum, start, end := maxSumSubArray(arr, k)
	fmt.Println("Sum is:", sum)
	fmt.Println("Sub array is:", arr[start:end+1])

	// 2. first negative integer in window of size k
	arr = []int{12, -1, -7, 8, -15, 30, 18, 28}
	k = 3
	ans := firstNegativeInt(arr, k)
	fmt.Println("Negative int in window of size k:", ans)

	// 3. count the occurences of anagram in a string
	s := "forxxorfxdofro"
	p := "for"
	countMap := make(map[string]int, 0)
	for _, char := range p {
		countMap[string(char)]++
	}
	occ := countOccurences(countMap, s, p)
	fmt.Println("No of occurences of anagram of p in s are:", occ)

	// 4. find the maximum of all sub array of size k
	arr = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = 3
	ans = maxOfSizeK(arr, k)
	fmt.Println("Max of all sub array of size k", ans)

	// Variable window size
	// 1. find largest sub array in the given array having sum (k)
	arr = []int{4, 1, 1, 1, 2, 3, 5}
	sum = 5
	start, end = largestSubArrayOfSumK(arr, sum)
	fmt.Println("Sub array having sum k is:", arr[start:end+1])

	// 2. longest sub string with k unique characters
	s = "aabacbebebe"
	start, end = longSubStrWithKUnique(s, 3)
	fmt.Println("Longest Sub String with k unique chars is:", s[start:end+1])

	// 3. longest sub string with all unique characters
	s = "pwwkerrr"
	start, end = longSubStringWithAllUnique(s)
	fmt.Println("Longest sub string without any repeating character is:", s[start:end+1])

	// 4. longest palindromic substring in a string
	s = "forgeeksskeegfor"
	start, end = palindromeSubStrOfMaxLen(s)
	fmt.Println("Longest Palindromic substring is:", s[start:end])
}
