package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	// for _, n := range nums {
	// 	total += n
	// }
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	all := append(a, b...)
	t := sum(all...)
	fmt.Println(t)
}
