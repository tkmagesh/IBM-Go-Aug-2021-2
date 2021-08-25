package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go add(2, 3, ch)
	result := <-ch //rading data from the channel
	wg.Wait()
	fmt.Println(result)

}

func add(x, y int, ch chan int) {
	result := x + y
	time.Sleep(3 * time.Second)
	ch <- result //writing data into the channel
	wg.Done()
}
