//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Character struct {
	health, max_health int
	energy, max_energy int
	name               string
}

func (c Character) Status() {
	fmt.Printf("Name: %v\nHealth: %v/%v\nEnergy: %v/%v\n", c.name, c.health, c.max_health, c.energy, c.max_energy)
}

func (c *Character) Damage(h int) {
	if c.health-h > 0 {
		c.health = c.health - h
	} else {
		c.health = 0
	}
	fmt.Println("Health changed to", c.health)
	c.Status()
}
func (c *Character) Effort(e int) {
	if c.energy-e > 0 {
		c.energy = c.energy - e
	} else {
		c.energy = 0
	}
	// fmt.Printf("Health changed to %d %d %d %d %s\n", c.health, c.max_health, c.energy, c.max_energy, c.name)
	c.Status()
	fmt.Println()
}
func main() {
	p1 := Character{health: 100, max_health: 100,
		energy: 100, max_energy: 100,
		name: "Haha"}

	fmt.Println(p1)
	p1.Damage(50)
	p1.Effort(29)

	p1.Damage(30)
	p1.Effort(10)
}
