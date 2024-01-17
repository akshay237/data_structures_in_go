package main

import (
	"fmt"
	"testing"
)

func TestCountOccurences(t *testing.T) {
	s := "forxxorfxdofro"
	p := "for"
	countMap := make(map[string]int, 0)
	for _, char := range p {
		countMap[string(char)]++
	}
	occ := countOccurences(countMap, s, p)
	fmt.Println("No of occurences of anagram of p in s are:", occ)
}
