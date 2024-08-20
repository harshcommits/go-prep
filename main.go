package main

import (
	"github.com/harshcommits/go-prep/concurrency"
)

func main() {
	// concurrency.DirectGoroutine()
	concurrency.RunWaitGroup()
	concurrency.SyncMutexFunc()
}
