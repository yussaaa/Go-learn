package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count++
	fmt.Println("Iteration", iteration, "Hits", hits.count)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	hitCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}
	fmt.Printf("Waiting for all tasks to complete\n")
	wg.Wait()

	hitCounter.Lock()
	fmt.Printf("All tasks done. Total hits: %d\n", hitCounter.count)
	hitCounter.Unlock()

}
