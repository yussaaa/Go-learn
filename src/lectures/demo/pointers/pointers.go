package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
	fmt.Println(counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	counter.hits++
	fmt.Println(counter.hits)
}

func main() {
	counter := Counter{}
	fmt.Println(counter.hits)

	hello := "Hello"
	world := "World"

	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	replace(&world, "Yusa", &counter)
	fmt.Println(hello, world)

	fmt.Println(counter.hits)
	increment(&counter)
	fmt.Println(counter.hits)
}
