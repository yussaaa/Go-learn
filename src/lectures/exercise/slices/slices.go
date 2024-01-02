//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printLine(line []Part) {
	fmt.Println("Assembly Line:")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}

func main() {
	// line := []Part{"part1", "part2", "part3"}
	line := make([]Part, 3)
	line[0] = "part1"
	line[1] = "part2"
	line[2] = "part3"
	printLine(line)

	line = append(line, "part4", "part5")
	printLine(line)
	printLine(line[3:])
}
