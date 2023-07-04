package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d: Starting task %d\n", id, task)
		time.Sleep(time.Second) // Simulating task execution time
		fmt.Printf("Worker %d: Completed task %d\n", id, task)
		wg.Done()
	}
}
func main() {
	numTasks := 15  // Number of tasks to be executed
	numWorkers := 5 // Number of workers
	tasks := make(chan int, numTasks)
	var wg sync.WaitGroup
	// Start the workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, &wg)
	}
	// Generate tasks
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		tasks <- i
	}
	close(tasks)
	startTime := time.Now()
	// Wait for all tasks to be completed
	wg.Wait()
	endTime := time.Now()
	// Calculate total time taken
	totalTime := endTime.Sub(startTime)
	fmt.Printf("Total time taken: %s\n", totalTime)
}
