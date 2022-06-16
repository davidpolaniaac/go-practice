package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Balance() int {
	b := balance
	return b
}

func deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}
func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}

//go build -o main --race mutex/mutex.go
