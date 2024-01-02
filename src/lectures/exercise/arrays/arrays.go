//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func printList(list [4]Product) {
	// fmt.Println(list[len(list)-1])
	// fmt.Println(len(list))
	// fmt.Println(list[0].price + list[1].price + list[2].price + list[3].price)
	var total, count int
	for i := 0; i < len(list); i++ {
		if list[i].name != "" {
			total += int(list[i].price * 100)
			count++
		}
	}
	fmt.Println(list[count-1])
	fmt.Println(count)
	fmt.Println(total)
}

func main() {
	shoppingList := [4]Product{
		{name: "milk", price: 2.99},
		{name: "eggs", price: 1.99},
		{name: "bread", price: 1.99},
	}
	printList(shoppingList)

	shoppingList[3] = Product{name: "cheese", price: 3.99}
	printList(shoppingList)
}
