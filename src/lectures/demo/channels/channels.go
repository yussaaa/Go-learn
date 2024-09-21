package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// One directional chanlle for jobs and result, job < - chan means only read from the channel, chan <- Result means only write to the channel
func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exiting goroutine!")
				control <- ExitOk
				return
			default:
				panic("Unhandeled control message")
			}
		case job := <-jobs:
			result := job.data * 2
			results <- Result{result, job}
		}
	}
}

func main() {
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
			control <- DoExit
			<-control // Reading from the control channel to wait for the goroutine to exit
			fmt.Println("Exiting program!")
			return
		}
	}

}
