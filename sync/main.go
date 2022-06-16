package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}
func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
