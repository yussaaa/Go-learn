//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func lineIterator(lines []string, line_func func(string)) {
	for _, line := range lines {
		// fmt.Println("\n", line)
		line_func(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letter_count := 0
	digit_count := 0
	space_count := 0
	punctuation_count := 0

	line_func := func(line string) {
		for _, r := range line {
			switch {
			case unicode.IsLetter(r):
				letter_count++
			case unicode.IsDigit(r):
				digit_count++
			case unicode.IsSpace(r):
				space_count++
			default:
				punctuation_count++
			}
		}

	}

	lineIterator(lines, line_func)
	fmt.Printf("Letters: %d\n", letter_count)
	fmt.Printf("Digits: %d\n", digit_count)
	fmt.Printf("Spaces: %d\n", space_count)
	fmt.Printf("Punctuation: %d\n", punctuation_count)
}
