package main

import "fmt"

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

func main() {

	// 1. print n num
	printNNum(10)

	// 2. factorial
	fact := factorial(5)
	fmt.Println("Factorial is:", fact)

	// 3. fibanocci
	fib := fibanocci(5)
	fmt.Println("Fibanocci is:", fib)
}
