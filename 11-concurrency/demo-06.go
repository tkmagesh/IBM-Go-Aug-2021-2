package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go add(2, 3, ch)
	result := <-ch //reading data from the channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	time.Sleep(3 * time.Second) //simulating a time consumering operation
	ch <- result                //writing data into the channel
}
