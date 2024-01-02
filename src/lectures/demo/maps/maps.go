package main

import "fmt"

func main() {
	shopping := make(map[string]int)

	shopping["shoes"] = 1
	shopping["socks"] = 4
	shopping["pants"] = 2

	shopping["socks"] += 1
	fmt.Println(shopping)

	delete(shopping, "pants")
	fmt.Println(shopping)

	for key, value := range shopping {
		fmt.Println(key, value)
	}

	socks_num, found := shopping["socks"]
	fmt.Println(socks_num, found)

	var totalItems int
	for _, value := range shopping {
		totalItems += value
	}
	fmt.Println(totalItems)
}
