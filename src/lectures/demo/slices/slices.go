package main

import "fmt"

func printSlice(name string, s []string) {
	fmt.Println("\n", "---", name, "---")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
func main() {
	route := []string{"home", "school", "store", "home"}

	printSlice("Route 1", route[:2])
	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")

	route = append(route, "park", "gym")

	printSlice("Route 2", route[2:])

}
