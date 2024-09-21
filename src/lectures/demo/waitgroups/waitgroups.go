package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter++
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter--
				wg.Done()

			}()
			duration := time.Duration(rand.Intn(500)) * time.Microsecond
			fmt.Println("Sleeping for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("Waiting for all tasks to complete")
	// wg.Wait()
	fmt.Println("All tasks done")
}
