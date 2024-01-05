package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// Normal function
func occupySpace(lot *ParkingLot, space int) {
	lot.spaces[space].occupied = true
}

// Receiver function
func (lot *ParkingLot) occupySpace(space int) {
	lot.spaces[space].occupied = true
}

func (lot *ParkingLot) vacateSpace(space int) {
	lot.spaces[space].occupied = false
}

func main() {
	// lot := [...]ParkingLot{
	// 	{[]Space{{false}, {false}, {false}, {false}, {false}}},
	// }
	lot := ParkingLot{spaces: make([]Space, 5)}
	fmt.Println(lot)

	lot.occupySpace(0)
	occupySpace(&lot, 1)
	fmt.Println(lot)

	lot.vacateSpace(0)
	fmt.Println(lot)

}
