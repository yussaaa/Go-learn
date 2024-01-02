//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) int {
	return r.length * r.width
}

func perimeter(r rectangle) int {
	return (r.length + r.width) * 2
}

func printresult(r rectangle) {
	fmt.Println("Area of rectangle is:", area(r))
	fmt.Println("Perimeter of rectangle is:", perimeter(r))
}

func main() {
	r1 := rectangle{length: 10, width: 5}
	printresult(r1)

	r1.length *= 2
	r1.width *= 2
	printresult(r1)

}
