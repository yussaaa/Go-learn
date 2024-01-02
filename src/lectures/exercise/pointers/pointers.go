//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type securityTag bool

const (
	Active   = true
	Inactive = false
)

type Item struct {
	name string
	tag  securityTag
}

func activate(tag *securityTag) {
	*tag = true
}
func deactivate(tag *securityTag) {
	*tag = false
}

// items[0] = securityTag{active: true}
// items[1] = securityTag{active: true}
// items[2] = securityTag{active: true}
// items[3] = securityTag{active: true}

func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		// item := items[i]
		deactivate(&items[i].tag)
	}
}

func main() {
	shirt := Item{name: "shirt", tag: true}
	shoes := Item{name: "shoes", tag: true}
	hat := Item{name: "hat", tag: true}
	pants := Item{"pants", Active}

	items := []Item{shirt, shoes, hat, pants}
	fmt.Println("Initial:", items)

	deactivate(&items[0].tag)
	fmt.Println("Item 0 Deactivated:", items)

	checkout(items)

	// for i := 0; i < len(items); i++ {
	// 	deactivate(&items[i].tag)
	// }
	fmt.Println(items)
}
