package main

import (
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go print("Hello", time.Second*2, ch1, ch2, wg)
	go print("World", time.Second*1, ch2, ch1, wg)
	ch1 <- 1
	wg.Wait()
}

func print(str string, delay time.Duration, ch1, ch2 chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-ch1
		time.Sleep(delay)
		println(str)
		ch2 <- 1
	}
	wg.Done()
}

/*
Desired Output
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
