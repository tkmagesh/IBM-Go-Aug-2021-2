package main

import (
	"fmt"
	"sync"
	"time"
)

//communicating by sharing memory - NOT ADVICABLE!!
var result int

var mutex sync.Mutex = sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Add(1)
	go add(1000, 2000, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	result = x + y
	time.Sleep(3 * time.Second)
	fmt.Println(result)
	mutex.Unlock()

	wg.Done()
}
