package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(l, r int) int {
	return l + r
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakerDozen := add(dozen, 1)
	fmt.Println("A bakerDezen is", bakerDozen)

	anotherBakerDozen := add(double(6), 1)
	fmt.Println("Another baker Dozen is", anotherBakerDozen)

}
