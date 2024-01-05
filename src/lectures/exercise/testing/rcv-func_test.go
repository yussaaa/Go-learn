// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestDamage(t *testing.T) {
	p1 := Character{health: 100, max_health: 100,
		energy: 100, max_energy: 100,
		name: "Haha"}

	p1.Damage(50)
	if p1.health > p1.max_health || p1.health < 0 {
		t.Errorf("Health exceed max or below 0")
	}
}

func TestEffort(t *testing.T) {
	p1 := Character{health: 100, max_health: 100,
		energy: 100, max_energy: 100,
		name: "Haha"}

	p1.Effort(50)
	if p1.energy > p1.max_energy || p1.energy < 0 {
		t.Errorf("Energy exceed max or below 0")
	}
}
