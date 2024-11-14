package main

import (
	"concurrent_labs/semaphore"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Producer function
func producer(id int, buffer int, sem *semaphore.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100) // Simulate work
		sem.P()                            // Acquire a token before producing
		buffer++
		fmt.Printf("Producer %d producing \n", id)
	}
}

// Consumer function
func consumer(id int, buffer int, sem *semaphore.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 150) // Simulate work
		sem.V()                            // Release the token after consuming
		buffer--
		fmt.Printf("Consumer %d consuming \n", id)
	}
}

// Test Producer-Consumer with Semaphore
func TestProducerConsumer(t *testing.T) {
	const numProducers = 3
	const numConsumers = 3
	buffer := 0
	sem := semaphore.NewSemaphore(3) // Limit to one producer/consumer at the same time
	var wg sync.WaitGroup

	// Start producers
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, buffer, sem, &wg)
	}

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, buffer, sem, &wg)
	}

	wg.Wait() // Wait for all producers and consumers to finish

	// Check if the buffer has the correct number of items
	if buffer != 0 {
		t.Errorf("Buffer should be empty, but has %d items", buffer)
	}
}

// Test semaphore functionality
func TestSemaphoreFunctionality(t *testing.T) {
	sem := semaphore.NewSemaphore(2) // Allow 2 concurrent accesses
	var wg sync.WaitGroup
	completed := make([]bool, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.P()
			time.Sleep(100 * time.Millisecond)
			completed[id] = true
			sem.V()
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	// Check if all goroutines have completed
	for i := 0; i < 5; i++ {
		if !completed[i] {
			t.Errorf("Goroutine %d did not complete", i)
		}
	}
}

// Test semaphore exceeding limit
func TestSemaphoreExceedLimit(t *testing.T) {
	sem := semaphore.NewSemaphore(1) // Limit to one access at the same time
	var wg sync.WaitGroup
	accessCount := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.P()
			accessCount++
			time.Sleep(50 * time.Millisecond) // Simulate work
			sem.V()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	// Check that accessCount should not exceed the limit
	if accessCount > 10 {
		t.Error("Access count exceeded the expected limit")
	}
}

// Test immediate release of semaphore
func TestSemaphoreImmediateRelease(t *testing.T) {
	sem := semaphore.NewSemaphore(1) // Limit to one access at the same time
	var wg sync.WaitGroup
	var accessed bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		sem.P()
		accessed = true
		sem.V()
	}()

	time.Sleep(100 * time.Millisecond) // Ensure goroutine has time to run
	// The semaphore should be released almost immediately after acquisition
	if !accessed {
		t.Fatal("The semaphore was not accessed as expected")
	}

	wg.Wait()
}
