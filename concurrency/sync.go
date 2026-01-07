package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// DemoMutex shows how to use sync.Mutex to protect shared resources
func DemoMutex() {
	var balance int
	var mutex sync.Mutex

	// Simulate concurrent deposits
	deposit := func(amount int, wg *sync.WaitGroup) {
		defer wg.Done()

		mutex.Lock()
		fmt.Printf("Depositing %d to account\n", amount)
		balance += amount
		mutex.Unlock()
	}

	var wg sync.WaitGroup
	// Start 5 concurrent deposits
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go deposit(i*100, &wg)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d\n", balance)
}

// DemoSyncMap shows how to use sync.Map for concurrent map operations
func DemoSyncMap() {
	var sm sync.Map

	// Concurrent writers
	for i := 0; i < 5; i++ {
		go func(n int) {
			key := fmt.Sprintf("key-%d", n)
			sm.Store(key, n*100)
			time.Sleep(100 * time.Millisecond)

			// Read our own write
			if val, ok := sm.Load(key); ok {
				fmt.Printf("Key %s has value: %v\n", key, val)
			}
		}(i)
	}

	// Concurrent reader
	go func() {
		sm.Range(func(key, value any) bool {
			fmt.Printf("Ranging: %v = %v\n", key, value)
			return true
		})
	}()

	// Give time for goroutines to complete
	time.Sleep(1 * time.Second)
}
