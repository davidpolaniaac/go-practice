package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkGroup(serve string, wg *sync.WaitGroup) {
	checkServe(serve)
	defer wg.Done()
}

func checkChannel(serve string, c chan<- string) {
	checkServe(serve)
	c <- serve
}

func checkSequential(serve string) {
	checkServe(serve)
}

func checkServe(serve string) {
	response, err := http.Get(serve)
	if err != nil {
		fmt.Println("Error", serve, response.Status)
	} else {
		fmt.Println("Successful", serve, response.Status)
	}
}

func waitGroup(servers []string) {
	var wg sync.WaitGroup
	for _, v := range servers {
		wg.Add(1)
		go checkGroup(v, &wg)
	}
	wg.Wait()
}

func Sequential(servers []string) {
	for _, v := range servers {
		checkSequential(v)
	}
}

func channel(servers []string) {
	channel := make(chan string)
	for _, v := range servers {
		go checkChannel(v, channel)
	}

	for i := 0; i < len(servers); i++ {
		<-channel
	}
}

func main() {

	servers := []string{
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/2",
	}

	startSequential := time.Now()
	Sequential(servers)
	fmt.Println("Time Sequential", time.Since(startSequential))

	startGroup := time.Now()
	waitGroup(servers)
	fmt.Println("Time Group", time.Since(startGroup))

	startChannel := time.Now()
	channel(servers)
	fmt.Println("Time Channel", time.Since(startChannel))

}
