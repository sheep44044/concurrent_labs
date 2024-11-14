package main

import (
	spinlock "concurrent_labs/spin_lock"
	"sync"
	"testing"
)

func TestSpinLock(t *testing.T) {
	lock := &spinlock.Spinlock{}
	var wg sync.WaitGroup
	const numGoroutines = 100
	const incrementsPerGoroutine = 1000
	var counter int32

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}

	wg.Wait()

	// Check if the counter value is as expected
	expected := int32(numGoroutines * incrementsPerGoroutine)
	if counter != expected {
		t.Errorf("expected counter to be %d, got %d", expected, counter)
	}
}
