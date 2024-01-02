package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{Name: "Casey", TicketNumber: 1, Boarded: true}
	bus := Bus{FrontSeat: casey}
	fmt.Println(bus.FrontSeat.Name)

	var (
		Bill  = Passenger{Name: "Bill", TicketNumber: 2, Boarded: false}
		Steve = Passenger{Name: "Steve", TicketNumber: 3, Boarded: false}
	)

	Bill.Boarded = true
	bus.FrontSeat = Steve
	fmt.Println(bus.FrontSeat.Name, bus.FrontSeat.TicketNumber, bus.FrontSeat.Boarded)
}
