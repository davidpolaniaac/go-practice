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

func deposit(amount int, wg *sync.WaitGroup) {
	b := balance
	balance = b + amount
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go deposit(i*100, &wg)
	}
	wg.Wait()
	fmt.Println(Balance())
}
