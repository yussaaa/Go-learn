package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{
		items:      make(map[P][]V),
		priorities: priorities,
	}
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for _, p := range pq.priorities {
		if len(pq.items[p]) > 0 {
			v := pq.items[p][0]
			pq.items[p] = pq.items[p][1:]
			return v, true
		}
	}
	var empty V
	return empty, false
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low}) // Here's where the priority is defined
	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())

}
