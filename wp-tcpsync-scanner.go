package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	// the worker(int, *sync.WaitGroup) func takes two arg:
	// a chan of type int and a pointer to a waitgroup
	// channel to recieve work, waitgroup to track when a single work is completed
	for p := range ports {
	fmt.Println(p)
	wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	// using make to create a channel, an int value of 100 is to make here
	var wg sync.WaitGroup
	
	for i:= 0; i<= cap(ports); i++ {
	// for loop to start the desired number of workers-pool 100; in this case
	// worker(int, *sync.WaitGroup) func: using range to recieve from ports channel
	// looping until the channel is closed
		go worker(ports, &wg)
	}
	for i:= 1; i<= 65535; i++ {
		wg.Add(1)
		
		ports <- 1
		// iterating over the ports sequentially in the main func
		// port is sent on the ports channel to the worker
	}
	wg.Wait()
	close(ports) // to close the ports channel
}

// we're using Worker Pool
// which uses a pool of goroutines to manage the concurrency
// still using WaitGroup block execution

