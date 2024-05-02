package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
	Water
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return [...]string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return [...]string{"Ground", "Air", "Water"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WareHouseAutomated interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}
func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WareHouseAutomated) {
	fmt.Printf("Conveying %v package of %v\n", item.Convey(), item)
	fmt.Printf("Shipping %v package of %v\n", item.Ship(), item)
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{amount: 100}
	automate(&mail)

	// waste := ToxicWaste{weight: 1000}
	// automate(&waste)
}
