package main

import "fmt"

type Room struct {
	name  string
	clean bool
}

func checkCleaness(r [4]Room) {
	for i := 0; i < len(r); i++ {
		if r[i].clean {
			fmt.Println(r[i].name, "is clean")
		} else {
			fmt.Println(r[i].name, "is dirty")
		}
	}

}

func main() {
	rooms := [4]Room{
		{name: "bedroom", clean: true},
		{name: "bathroom", clean: false},
		{name: "kitchen", clean: true},
		{name: "livingroom", clean: false},
	}

	checkCleaness(rooms)
	rooms[1].clean = true
	rooms[3].clean = true

	checkCleaness(rooms)

}
