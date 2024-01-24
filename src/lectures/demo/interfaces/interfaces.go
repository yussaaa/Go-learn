package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cooking chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Mix salad")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Preparing dishes")

	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("-- Preparing dish %s --\n", dish)
		dish.PrepareDish()
		fmt.Println("")
	}
}

func main() {
	dishes := []Preparer{Chicken("Chicken"), Salad("Salad")}

	prepareDishes(dishes)
}
