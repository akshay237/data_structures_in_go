package main

import (
	"fmt"
	"testing"
)

func TestSortUsingRecursion(t *testing.T) {
	arr := []int{1, 5, 7, 2, 4, 3}
	newarr := sortUsingRecursion(arr)
	fmt.Println("Arr", newarr)
}

func TestReverseStack(t *testing.T) {
	arr := []int{1, 5, 7, 2, 4, 3}
	newarr := reverseStackUsingRecursion(arr)
	fmt.Println("Reversed stack is", newarr)
}
