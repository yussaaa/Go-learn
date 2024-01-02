package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, v := range slice {
		fmt.Println(i, v)
		for _, ch := range v {
			fmt.Println("  ", string(ch))
		}
	}
}
