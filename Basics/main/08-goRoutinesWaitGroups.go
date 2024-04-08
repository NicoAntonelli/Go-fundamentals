package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Concurrency with wait groups
var waitGroup = sync.WaitGroup{}

// Mutex for safe concurrent access
var mutex = sync.Mutex{}
var mutexRW = sync.RWMutex{}

// Mocked DB data
var dbData = []string{"alpha", "beta", "gamma", "delta", "epsilon"}
var results = []string{}

// Work with Go routines using wait groups & mutex
func goRoutinesWaitGroups() {
	fmt.Println("Go Routines!")
	timeBase := time.Now()
	for i := 0; i < len(dbData); i++ {
		// Increase wait group counter by 1
		waitGroup.Add(1)
		// Go routine is async (concurrency)
		go dbCallMock(i)
	}
	// Await until wait group drops to zero
	waitGroup.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(timeBase))
	fmt.Printf("The results are: %v\n", results)
	fmt.Println()
}

// Simulate DB call delay
func dbCallMock(id int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The %v° result from the DB is: %v\n", id+1, dbData[id])
	
	// Append results safer with mutex locking
	mutex.Lock()
	results = append(results, dbData[id])
	mutex.Unlock()

	// Shows results with RWMutex locking
	mutexRW.RLock()
	fmt.Printf("The current results with %v° are: %v\n", id+1, results)
	mutexRW.RUnlock()

	// Decrease wait group counter by 1
	waitGroup.Done()
}
