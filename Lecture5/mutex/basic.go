package main

import (
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter = counter + 1
		}()
	}
	time.Sleep(1 * time.Second)
	mu.Lock()
	// does not guarantee to print 1000
	println(counter)
	mu.Unlock()
}
