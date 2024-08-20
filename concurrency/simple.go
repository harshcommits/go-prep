package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Hello() {
	fmt.Println("Hello World")
}

func Goodbye() {
	fmt.Println("Goodbye")
}

// same hello function as previous implemented with waitgroup
func Hellowg(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}

func DirectGoroutine() {
	go Hello()
	time.Sleep(1 * time.Second)
	Goodbye()
}

func RunWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Hellowg(&wg)
	wg.Wait()
	Goodbye()

}

func SyncMutexFunc() {
	// Create a new sync.Map
	var m sync.Map

	// Number of goroutines for concurrent operations
	numGoroutines := 5

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Writing goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Store key-value pair
			m.Store(id, id)

			// Load and print the value
			if value, ok := m.Load(id); ok {
				fmt.Printf("Goroutine %d: Key %d - Value %d\n", id, id, value)
			}
		}(i)
	}

	// Wait for all writing goroutines to finish
	wg.Wait()

	// Reading goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// Load and print the value
			if value, ok := m.Load(id); ok {
				fmt.Printf("Reading Goroutine %d: Key %d - Value %d\n", id, id, value)
			} else {
				fmt.Printf("Reading Goroutine %d: Key %d not found\n", id, id)
			}
		}(i)
	}

	// Wait for all reading goroutines to finish
	wg.Wait()
}
