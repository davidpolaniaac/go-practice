package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) createDatabaseInstance() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(5 * time.Second)
}

var database *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if database == nil {
		fmt.Println("Creating DB Connection")
		database = &Database{}
		database.createDatabaseInstance()
	} else {
		fmt.Println("DB Already Created")
	}
	return database
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
