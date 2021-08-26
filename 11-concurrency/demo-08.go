package main

import (
	"fmt"
	"sync"
	"time"
)

//var wg *sync.WaitGroup = &sync.WaitGroup{}
var wg sync.WaitGroup

func main() {
	count := make(chan int)
	wg.Add(1)
	go fn(count)
	time.Sleep(time.Second * 4)
	count <- 10 // channel used to send a signal to the goroutine
	wg.Wait()
	fmt.Println("[main] exiting")
}

func fn(count chan int) {
	fmt.Println("[fn] Attempting to read the count")
	max := <-count
	fmt.Println("[fn] The count is", max)
	for i := 0; i < max; i++ {
		fmt.Println("[fn] - ", i)
	}
	wg.Done()
}
