package main

import "fmt"

func main() {
	var myName = "Yusa"
	fmt.Println("My name is:", myName)

	var name string = "Yusa"
	fmt.Println("Name:", name)

	userName := "Yusa"
	fmt.Println(userName)

	// Default value on int is 0
	var sum int
	fmt.Println(sum)

	part1, other := 1, 2
	fmt.Println(part1, other)

	var (
		lessonName = "Variable"
		lessonType = "Demo"
	)
	fmt.Println(lessonName, lessonType)

}
