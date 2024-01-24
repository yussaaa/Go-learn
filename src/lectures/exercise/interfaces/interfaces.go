//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type Director interface {
	Direct() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) Direct() Lift {
	// fmt.Println("Go to small lift")
	return SmallLift
}

func (c Car) Direct() Lift {
	// fmt.Println("Go to standard lift")
	return StandardLift
}

func (t Truck) Direct() Lift {
	// fmt.Println("Go to large lift")
	return LargeLift
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func directVehicles(vehicles []Director) {
	fmt.Println("Directing vehicles")

	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		fmt.Printf("-- Directing vehicle %s --\n", vehicle)
		switch vehicle.Direct() {
		case SmallLift:
			fmt.Println("Go to small lift")
		case StandardLift:
			fmt.Println("Go to standard lift")
		case LargeLift:
			fmt.Printf("%v Go to large lift", vehicle)
		}
		fmt.Println("")
	}
}

func main() {
	vehicles := []Director{Motorcycle("Motorcycle"), Car("Car"), Truck("Tesla-Truck")}

	directVehicles(vehicles)

}
