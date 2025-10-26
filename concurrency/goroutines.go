package concurrency

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello world!")
}

func goodbye() {
	fmt.Println("Goodbye world!")
}

func Goroutine() {

	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	// time.Sleep(1 * time.Second)
	wg.Wait()
	goodbye()
}
