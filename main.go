package main

import (
	"fmt"
	"sync"
	"time"
)

func scheduleTaskOnce(duration time.Duration, task func()) {
	// Use time.After to create a channel that will be sent the
	// current time after the specified duration.
	timer := time.After(duration)

	// Use a select statement to wait for the timer channel to be sent
	// the current time.
	select {
	case <-timer:
		// When the timer channel is sent the current time, call the
		// task function.
		go task()
	}
}

func schedular(wg *sync.WaitGroup, duration time.Duration, task func()) {
	// Add a count to the wait group before starting the goroutine.
	wg.Add(1)
	go func() {
		// Schedule the task to be executed after the specified duration.
		scheduleTaskOnce(duration, task)
		// Decrement the wait group count when the goroutine finishes.
		defer wg.Done()
	}()
}

func main() {
	fmt.Println("start")

	// Create a new wait group.
	var wg sync.WaitGroup

	// Schedule the tasks to be executed at different times.
	schedular(&wg, 5*time.Second, func() {
		fmt.Println("5 seconds have elapsed")
	})
	schedular(&wg, 8*time.Second, func() {
		fmt.Println("8 seconds have elapsed")
	})
	schedular(&wg, 11*time.Second, func() {
		fmt.Println("11 seconds have elapsed")
	})

	fmt.Println("before sleep")
	time.Sleep(1 * time.Minute)
	fmt.Println("after sleep")
}
