package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func compute(a, b int, fn func(a, b int) int) int {
	fmt.Printf("Running a computtation with %v & %v\n", a, b)
	return fn(3, 4)
}

func main() {
	fmt.Println("3 + 4 = ", compute(3, 4, add))
	fmt.Println("3 - 4 = ", compute(3, 4, func(a, b int) int {
		return a - b
	}))

	mul := func(a, b int) int {
		fmt.Println("Running a multiplication with", a, b)
		return a * b
	}
	fmt.Println("3 * 4 = ", compute(3, 4, mul))
}
