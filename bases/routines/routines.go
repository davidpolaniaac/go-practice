package routines

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)
}

func Run() {

	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1)

	go say("world", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Bye")

	time.Sleep(time.Second * 1)

}
