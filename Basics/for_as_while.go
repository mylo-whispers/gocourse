package Basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Retry pattern with exponential backoff
	maxRetries := 5
	for retries := 0; retries < maxRetries; retries++ {
		fmt.Printf("Attempt %d: ", retries+1)

		// Simulate some operation that might fail
		success := rand.Float32() < 0.3 // 30% success rate

		if success {
			fmt.Println("Success!")
			break
		} else {
			fmt.Println("Failed, retrying...")
			backoff := time.Duration(retries*retries) * 100 * time.Millisecond
			time.Sleep(backoff)
		}

		if retries == maxRetries-1 {
			fmt.Println("Max retries reached, giving up")
		}
	}

	// 2. Processing until empty
	queue := []string{"task1", "task2", "task3", "task4"}
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		fmt.Printf("Processing: %s, remaining: %d\n", task, len(queue))

		// Simulate adding new tasks sometimes
		if rand.Float32() < 0.3 && len(queue) < 10 {
			newTask := fmt.Sprintf("new-task-%d", len(queue))
			queue = append(queue, newTask)
			fmt.Printf("Added new task: %s\n", newTask)
		}
	}

	// 3. Loop with deferred cleanup
	for i := 0; i < 3; i++ {
		fmt.Printf("Starting iteration %d\n", i)

		// Deferred function will execute at the end of each iteration
		defer func(iter int) {
			fmt.Printf("Cleaning up iteration %d\n", iter)
		}(i)

		// Simulate work
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Finished work in iteration %d\n", i)
	}
}
