package main

import (
	"fmt"
	"sync"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)   // sleep for 5 seconds
	fmt.Println("Done Goroutine") // We won't see this
}

func main() {
	fmt.Println("Start")
	heavy() // Normal execution
	fmt.Println("End")

	fmt.Println("Start Goroutine")
	go heavy() // Goroutine execution
	fmt.Println("End Goroutine")

	// WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1) // Add 1 to the WaitGroup (the following goroutine)
	go func() {
		heavy()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("End WaitGroup")

	// Mutex
	mu := &sync.Mutex{}
	mu.Lock()
	go func() {
		heavy()
		mu.Unlock()
	}()
	mu.Lock()
}
