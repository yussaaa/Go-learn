package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum:", sum)

	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	for sum > 10 {
		sum -= 10
		fmt.Println("Decrement sum is:", sum)
	}
}
