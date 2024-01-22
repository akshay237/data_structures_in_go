package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
)

func printNNum(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNNum(n - 1)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibanocci(n int) int {
	if n < 2 {
		return 1
	}
	return fibanocci(n-1) + fibanocci(n-2)
}

func insert(arr []int, temp int) []int {
	if len(arr) == 0 || arr[len(arr)-1] <= temp {
		arr = append(arr, temp)
		return arr
	}
	idx := len(arr) - 1
	val := arr[idx]
	arr = slices.Delete(arr, idx, idx+1)
	arr = insert(arr, temp)
	arr = append(arr, val)
	return arr
}

func sortUsingRecursion(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	idx := len(arr) - 1
	temp := arr[idx]
	arr = slices.Delete(arr, idx, idx+1)
	sortUsingRecursion(arr)
	return insert(arr, temp)
}

func deleteStackMiddleElement(stack []int, k int) []int {
	if k == 0 {
		idx := len(stack) - 1
		stack = slices.Delete(stack, idx, idx+1)
		return stack
	}
	index := len(stack) - 1
	temp := stack[index]
	stack = slices.Delete(stack, index, index+1)
	stack = deleteStackMiddleElement(stack, k-1)
	stack = append(stack, temp)
	return stack
}

func insertReverse(arr []int, val int) []int {
	if len(arr) == 0 {
		arr = append(arr, val)
		return arr
	}
	idx := len(arr) - 1
	temp := arr[idx]
	arr = slices.Delete(arr, idx, idx+1)
	arr = insertReverse(arr, val)
	arr = append(arr, temp)
	return arr
}

func reverseStackUsingRecursion(stack []int) []int {
	if len(stack) == 0 {
		return []int{}
	}
	idx := len(stack) - 1
	val := stack[idx]
	stack = slices.Delete(stack, idx, idx+1)
	stack = reverseStackUsingRecursion(stack)
	return insertReverse(stack, val)
}

func kthGrammar(n, k int) int {
	if n == 1 || k == 1 {
		return 0
	}

	length := int(math.Pow(float64(2), float64(n))) - 1
	mid := length / 2

	if k <= mid {
		return kthGrammar(n-1, k)
	} else {
		return ^kthGrammar(n-1, mid-k)
	}
}

func towerOfHanoi(S, D, H string, n int) {
	if n == 1 {
		fmt.Println("Move plate 1 from", S, "to", D, "with help of", H)
		return
	}
	towerOfHanoi(S, H, D, n-1)
	fmt.Println("Move plate", n, "from", S, "to", D, "with help of", H)
	towerOfHanoi(H, D, S, n-1)
}

// Problems based on the recursion tree
func printSubsetsOfString(input, op string, ans []string) []string {
	if len(input) == 0 {
		if !slices.Contains(ans, op) {
			ans = append(ans, op)
		}
		return ans
	}
	op1 := op
	op2 := op
	op2 += string(input[0])

	ans = printSubsetsOfString(input[1:], op1, ans) // either it is a subset or not
	ans = printSubsetsOfString(input[1:], op2, ans)
	return ans
}

func permutationWithSpaces(input, output string, ans []string) []string {
	if len(input) == 0 {
		ans = append(ans, output)
		return ans
	}

	// without spaces
	op1 := output
	op1 += string(input[0])
	// with spaces
	op2 := output
	op2 += "_"
	op2 += string(input[0])

	ans = permutationWithSpaces(input[1:], op1, ans)
	ans = permutationWithSpaces(input[1:], op2, ans)
	return ans
}

func permutationWithCaseChange(input, output string, ans []string) []string {
	if len(input) == 0 {
		ans = append(ans, output)
		return ans
	}

	// without case change
	output1 := output
	output1 += string(input[0])

	// with case change
	output2 := output
	output2 += strings.ToUpper(string(input[0]))

	// call the func
	ans = permutationWithCaseChange(input[1:], output1, ans)
	ans = permutationWithCaseChange(input[1:], output2, ans)
	return ans

}

func letterCaseChange(input, output string, ans []string) []string {
	if len(input) == 0 {
		ans = append(ans, output)
		return ans
	}

	// if character is alpha
	if unicode.IsLetter(rune(input[0])) {
		var ch string
		if unicode.IsLower(rune(input[0])) {
			ch = strings.ToUpper(string(input[0]))
		} else {
			ch = strings.ToLower(string(input[0]))
		}
		// swap the case
		op1 := output
		op1 += ch
		// with out swapping
		op2 := output
		op2 += string(input[0])
		// call for op1 and op2
		ans = letterCaseChange(input[1:], op1, ans)
		ans = letterCaseChange(input[1:], op2, ans)

	} else {
		op3 := output
		op3 += string(input[0])
		ans = letterCaseChange(input[1:], op3, ans)
	}
	return ans
}

func generateParenthesis(open, close int, op string, parentparentesis []string) []string {

	if open == 0 && close == 0 {
		parentparentesis = append(parentparentesis, op)
		return parentparentesis
	}

	if open > 0 {
		op1 := op
		op1 += "("
		parentparentesis = generateParenthesis(open-1, close, op1, parentparentesis)
	}

	if close > open {
		op2 := op
		op2 += ")"
		parentparentesis = generateParenthesis(open, close-1, op2, parentparentesis)
	}

	return parentparentesis
}

func printNBitBinaryNum(n, o, z int, op string, ans []string) []string {
	if n == 0 {
		ans = append(ans, op)
		return ans
	}

	if o > z {
		op1 := op
		op1 += "0"
		ans = printNBitBinaryNum(n-1, o, z+1, op1, ans)
	}

	op2 := op
	op2 += "1"
	ans = printNBitBinaryNum(n-1, o+1, z, op2, ans)

	return ans
}

func main() {

	// 1. print n num
	printNNum(10)

	// 2. factorial
	fact := factorial(5)
	fmt.Println("Factorial is:", fact)

	// 3. fibanocci
	fib := fibanocci(5)
	fmt.Println("Fibanocci is:", fib)

	// 4. sort an array using recursion
	arr := []int{1, 5, 7, 2, 4, 3}
	arr = sortUsingRecursion(arr)
	fmt.Println("Array after sorting is:", arr)

	// 5. delete middle element of stack
	arr = []int{1, 2, 4, 5, 7}
	middle := len(arr) / 2
	arr = deleteStackMiddleElement(arr, middle)
	fmt.Println("Stack after deleting middle element", arr)

	// 6. reverse a stack using recursion
	arr = []int{1, 2, 4, 5, 7}
	arr = reverseStackUsingRecursion(arr)
	fmt.Println("Reversed stack is:", arr)

	/* 7. We have to return the kth symbol from the nth row in the generated grammar with the given constrained  such as when n = 1 and
	k =1 return 0 and 0 --> 0 1 and 1 --> 1 0
	n = 2 depends n -1 so --> 0 1
	n = 3 depends on n - 2 --> 0 1 1 0
	n = 4 depends on n - 3 --> 0 1 1 0 1 0 0 1
	output: 0 or 1
	*/
	gram := kthGrammar(4, 7)
	if gram == 1 {
		fmt.Println("Grammar value is 1")
	} else {
		fmt.Println("Grammar value is 0")
	}

	// 8. Tower of Hanoi
	towerOfHanoi("A", "B", "C", 3)

	// 9. print subsets of a string
	subsets := []string{}
	op := " "
	subsets = printSubsetsOfString("aab", op, subsets)
	fmt.Println("Subsets of string are:", subsets)

	// 10. Permutation of a string with spaces
	s := "ABC"
	op = string(s[0])
	perms := []string{}
	perms = permutationWithSpaces(s[1:], op, perms)
	fmt.Println("Permutations of string with spaces are", perms)

	// 11. Permutations with case change
	s = "abc"
	op = ""
	permsCase := []string{}
	permsCase = permutationWithCaseChange(s, op, permsCase)
	fmt.Println("Permutations with case changes are", permsCase)

	// 12. Letter Case Change
	s = "a1B2"
	op = ""
	caseChangedOp := []string{}
	caseChangedOp = letterCaseChange(s, op, caseChangedOp)
	fmt.Println("Case changed op is", caseChangedOp)

	// 13. Generate Balanced Parenthesis
	n := 3
	op = ""
	open := n
	close := n
	parentesis := []string{}
	parentesis = generateParenthesis(open, close, op, parentesis)
	fmt.Println("Generated Balanced Parenthesis are", parentesis)

	// 14. Print a N bit binary number having more no of 1's than 0's in all of it's prefixes
	n = 5
	op = ""
	o := 0
	z := 0
	nums := []string{}
	nums = printNBitBinaryNum(n, o, z, op, nums)
	fmt.Println("N Bit binary nums are", nums)
}
