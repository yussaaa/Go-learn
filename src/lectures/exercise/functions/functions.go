//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func bye() {
	fmt.Println("Bye, hope you enjoied using this function!")
}

func add(a, b, c int) int {
	output := a + b + c
	return output
}

func multiply(x int) (int, int) {
	return 2 * x, 5 * x
}

func main() {
	greet("Yusa")

	add123 := add(1, 2, 3)
	fmt.Println("Sum of 1, 2, 3 is", add123)

	m, n := multiply(2)
	fmt.Println("Multiply 2 by 2 and 5 seperately is", m, n)

	bye()

}
