package main

import (
	"fmt"
	"sort"
)

func ReverseString(str string) string {
	rstr := ""
	for i := len(str) - 1; i >= 0; i -= 1 {
		rstr += string(str[i])
	}
	return rstr
}

func ReverseUsingRune(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i <= j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(str string) bool {
	revStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	return str == revStr
}

func isAnagrams(s1, s2 string) bool {

	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		return false
	}

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	s1Sorted := string(r1)
	s2Sorted := string(r2)

	return s1Sorted == s2Sorted
}

func main() {

	// 1. reverse a string using extra word
	word := "akshay"
	revWord := ReverseString(word)
	fmt.Println("Reversed string is: ", revWord)

	// 2. reverse string using rune
	revWordUsingRune := ReverseUsingRune(word)
	fmt.Println("Reversed string using rune: ", revWordUsingRune)

	// 3. string is palindrome
	isPali := isPalindrome(word)
	fmt.Println(word, "is a plaindrome ", isPali)

	s := "madam"
	isPali = isPalindrome(s)
	fmt.Println(s, "is a plaindrome ", isPali)

	// 4. Check if two strings are anagram or not
	s1 := "anagram"
	s2 := "ramanag"
	isok := isAnagrams(s1, s2)
	fmt.Println(s1, s2, "are anagrams:", isok)

}
